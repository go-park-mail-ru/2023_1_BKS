package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/jwt"
)

const PrivateKey = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIN2dALnjdcZaIZg4QuA6Dw+kxiSW502kJfmBN3priIhPoAoGCCqGSM49
AwEHoUQDQgAE4pPyvrB9ghqkT1Llk0A42lixkugFd/TBdOp6wf69O9Nndnp4+HcR
s9SlG/8hjB2Hz42v4p3haKWv3uS1C6ahCQ==
-----END EC PRIVATE KEY-----`

const KeyID = `fake-key-id`
const FakeIssuer = "fake-issuer"
const FakeAudience = "example-users"
const PermissionsClaim = "perm"

type JWSValidator interface {
	ValidateJWS(jws string) (jwt.Token, error)
}

var (
	ErrNoAuthHeader      = errors.New("Authorization header is missing")
	ErrInvalidAuthHeader = errors.New("Authorization header is malformed")
	ErrClaimsInvalid     = errors.New("Provided claims do not match expected scopes")
)

func GetJWSFromRequest(req *http.Request) (string, error) {
	authHdr := req.Header.Get("Authorization")

	if authHdr == "" {
		return "", ErrNoAuthHeader
	}

	prefix := "Bearer "
	if !strings.HasPrefix(authHdr, prefix) {
		return "", ErrInvalidAuthHeader
	}
	return strings.TrimPrefix(authHdr, prefix), nil
}

func NewAuthenticator(v JWSValidator) openapi3filter.AuthenticationFunc {
	return func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
		return Authenticate(v, ctx, input)
	}
}

func Authenticate(v JWSValidator, ctx context.Context, input *openapi3filter.AuthenticationInput) error {
	if input.SecuritySchemeName != "BearerAuth" {
		return fmt.Errorf("security scheme %s != 'BearerAuth'", input.SecuritySchemeName)
	}

	jws, err := GetJWSFromRequest(input.RequestValidationInput.Request)
	if err != nil {
		return fmt.Errorf("getting jws: %w", err)
	}

	token, err := v.ValidateJWS(jws)
	if err != nil {
		return fmt.Errorf("validating JWS: %w", err)
	}

	err = CheckTokenClaims(input.Scopes, token)

	if err != nil {
		return fmt.Errorf("token claims don't match: %w", err)
	}
	return nil
}

func GetClaimsFromToken(t jwt.Token) ([]string, error) {
	rawPerms, found := t.Get(PermissionsClaim)
	if !found {
		return make([]string, 0), nil
	}

	rawList, ok := rawPerms.([]interface{})
	if !ok {
		return nil, fmt.Errorf("'%s' claim is unexpected type'", PermissionsClaim)
	}

	claims := make([]string, len(rawList))

	for i, rawClaim := range rawList {
		var ok bool
		claims[i], ok = rawClaim.(string)
		if !ok {
			return nil, fmt.Errorf("%s[%d] is not a string", PermissionsClaim, i)
		}
	}
	return claims, nil
}

func CheckTokenClaims(expectedClaims []string, t jwt.Token) error {
	claims, err := GetClaimsFromToken(t)
	if err != nil {
		return fmt.Errorf("getting claims from token: %w", err)
	}
	claimsMap := make(map[string]bool, len(claims))
	for _, c := range claims {
		claimsMap[c] = true
	}

	for _, e := range expectedClaims {
		if !claimsMap[e] {
			return ErrClaimsInvalid
		}
	}
	return nil
}

func CreateMiddleware(v JWSValidator, swagger *openapi3.T) ([]echo.MiddlewareFunc, error) {

	validator := oapimiddleware.OapiRequestValidatorWithOptions(swagger,
		&oapimiddleware.Options{
			Options: openapi3filter.Options{
				AuthenticationFunc: NewAuthenticator(v),
			},
		})

	return []echo.MiddlewareFunc{validator}, nil
}
