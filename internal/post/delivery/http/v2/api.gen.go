// Package v2 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package v2

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Закрыть объявление.
	// (PUT /api/close/{userId}/{id})
	ClosePost(ctx echo.Context, userId string, id string) error
	// Создать новое объявление.
	// (POST /api/post)
	CreatePost(ctx echo.Context) error
	// Вернуть закрытые объявления по id пользователя.
	// (GET /api/post/close/user/{idUser}/{page})
	FindClosePostByUserID(ctx echo.Context, idUser string, page int) error
	// Вернуть открытые объявления по id пользователя.
	// (GET /api/post/open/user/{idUser}/{page})
	FindOpenPostByUserID(ctx echo.Context, idUser string, page int) error
	// Удалить объявление.
	// (DELETE /api/post/{id})
	DeletePost(ctx echo.Context, id string) error
	// Вернуть объявление по id.
	// (GET /api/post/{id})
	FindPostByID(ctx echo.Context, id string) error
	// Обновить объявление.
	// (PUT /api/post/{id})
	UpdatePost(ctx echo.Context, id string) error
	// Вернуть объявления по тегу.
	// (GET /api/post/{tag}/{page})
	FindPostByTag(ctx echo.Context, tag string, page int) error
	// Основнвая страница
	// (GET /api/sort/new/{page})
	GetAllPost(ctx echo.Context, page int) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// ClosePost converts echo context to params.
func (w *ServerInterfaceWrapper) ClosePost(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "userId", runtime.ParamLocationPath, ctx.Param("userId"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ClosePost(ctx, userId, id)
	return err
}

// CreatePost converts echo context to params.
func (w *ServerInterfaceWrapper) CreatePost(ctx echo.Context) error {
	var err error

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreatePost(ctx)
	return err
}

// FindClosePostByUserID converts echo context to params.
func (w *ServerInterfaceWrapper) FindClosePostByUserID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "idUser" -------------
	var idUser string

	err = runtime.BindStyledParameterWithLocation("simple", false, "idUser", runtime.ParamLocationPath, ctx.Param("idUser"), &idUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter idUser: %s", err))
	}

	// ------------- Path parameter "page" -------------
	var page int

	err = runtime.BindStyledParameterWithLocation("simple", false, "page", runtime.ParamLocationPath, ctx.Param("page"), &page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindClosePostByUserID(ctx, idUser, page)
	return err
}

// FindOpenPostByUserID converts echo context to params.
func (w *ServerInterfaceWrapper) FindOpenPostByUserID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "idUser" -------------
	var idUser string

	err = runtime.BindStyledParameterWithLocation("simple", false, "idUser", runtime.ParamLocationPath, ctx.Param("idUser"), &idUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter idUser: %s", err))
	}

	// ------------- Path parameter "page" -------------
	var page int

	err = runtime.BindStyledParameterWithLocation("simple", false, "page", runtime.ParamLocationPath, ctx.Param("page"), &page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindOpenPostByUserID(ctx, idUser, page)
	return err
}

// DeletePost converts echo context to params.
func (w *ServerInterfaceWrapper) DeletePost(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeletePost(ctx, id)
	return err
}

// FindPostByID converts echo context to params.
func (w *ServerInterfaceWrapper) FindPostByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindPostByID(ctx, id)
	return err
}

// UpdatePost converts echo context to params.
func (w *ServerInterfaceWrapper) UpdatePost(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	ctx.Set(BearerAuthScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdatePost(ctx, id)
	return err
}

// FindPostByTag converts echo context to params.
func (w *ServerInterfaceWrapper) FindPostByTag(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "tag" -------------
	var tag string

	err = runtime.BindStyledParameterWithLocation("simple", false, "tag", runtime.ParamLocationPath, ctx.Param("tag"), &tag)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tag: %s", err))
	}

	// ------------- Path parameter "page" -------------
	var page int

	err = runtime.BindStyledParameterWithLocation("simple", false, "page", runtime.ParamLocationPath, ctx.Param("page"), &page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindPostByTag(ctx, tag, page)
	return err
}

// GetAllPost converts echo context to params.
func (w *ServerInterfaceWrapper) GetAllPost(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "page" -------------
	var page int

	err = runtime.BindStyledParameterWithLocation("simple", false, "page", runtime.ParamLocationPath, ctx.Param("page"), &page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter page: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAllPost(ctx, page)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.PUT(baseURL+"/api/close/:userId/:id", wrapper.ClosePost)
	router.POST(baseURL+"/api/post", wrapper.CreatePost)
	router.GET(baseURL+"/api/post/close/user/:idUser/:page", wrapper.FindClosePostByUserID)
	router.GET(baseURL+"/api/post/open/user/:idUser/:page", wrapper.FindOpenPostByUserID)
	router.DELETE(baseURL+"/api/post/:id", wrapper.DeletePost)
	router.GET(baseURL+"/api/post/:id", wrapper.FindPostByID)
	router.PUT(baseURL+"/api/post/:id", wrapper.UpdatePost)
	router.GET(baseURL+"/api/post/:tag/:page", wrapper.FindPostByTag)
	router.GET(baseURL+"/api/sort/new/:page", wrapper.GetAllPost)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xZT2/bxhP9KsT+fkdCVOKedHPipnWBIgbqtIfUh420lragSGZ35cAQCEgykAZ1ACO9",
	"tCiQBGmBnmXFipUoZr7C7Dcqdpb6T1kK4tQ2qpPFf7Mzs++9eaTrpBhWozBggZKkUCeyWGFVij9vC0YV",
	"2wqlMkfU9+/uksL9OolEGDGhOMO7NpgsCh4pHgbmsDR+SOAVJHACXd2AN9CGM+hB14EEjvUv+gg60Ieu",
	"OamPcsQlaj9ipECkEjwok9glW1RVNqu0bBfiilVlxhK/wylGbEAb3th40M2Ml56gQtB9jC94kWVE/Buj",
	"tJdPdJuWM8L8AW3dgi68hkQ35j/LlZ+VxHNowyl0bNf00bLJxC4R7GGNC1YihftpdHdilwaF27wn2rwT",
	"78Qu+VKIUHy9vb1l0prc7WJYykgWH3Dwmkt2Q1GlihQID9TazVGGPFCszISpucqkpOW5gQaXFxWXLji4",
	"3aT+LQ/4IsR+dliFUm2WZkNubnwE9C8MmhcLL5fck0zMq+4DJNDXT7FznRT9/TlxvufskcxkTQJ96Omf",
	"oaubugUdSBz4oBuQ6Ca8h0S3zG/o5DKQNYWPNFV3yIJ0ZwaLj4gwQ4FFGLrthzKrrb9BG97phj7ULZN3",
	"H3oZnZ3A0YMw9BkNTEtWUnqlpPQcrJeuPtbPU3yL3jEaTMI/dolkxZrgav87YwcsaB4wKphYr6nK6OjO",
	"QOu/+WGbuNY8IKjx6ijrilIRiU1gHuyGs524G7FgPeLOWi7vwInpo6ObiPUO9HQT2k4USmXOteGtbkHb",
	"WY+iewF/iJ2x207SU86PtXx+rSiZ2ONF5hgm4xnmrG9tEpfsMSHtqjdy+Vze7E8YsYBGnBTIGp5ySURV",
	"Bav2aMS9oumXV69hj2OvzksxjsYaaoTRBWoKMVCxvUX5MFEErTLFhEQJWXoccHPdpEBcElDsKDd7O9pw",
	"JWos7Tc1OcwMyjnynA3YjPVqAzwtv+aOuVlGYSAtYG7m89YyBIoFVk2jyOdF7JX3k7QqN4r3f8F2SYH8",
	"zxs5Ui+1ox42FAE0RaEXs+Lq6APdhA/Q1U/gzDAqgWPzY3hPkiMYaJfWfHVhKY5cU1aeL5HJPdN9/QT6",
	"KIn6CfTgGN5BO0fGSYdgGafb/R3TXFmrVqnYnxoz+umcCWMiInqjdJYN/k7hdeTy7WYzqW6Fpf0L68vY",
	"AvGkahlAxTOguZEhlEvssm5CAqdwgip//Xf41aAau8MWvknmyJ/Z7FSvDIWNVpnREHv1iJYZylaZZcDg",
	"Dg9KQ+m6tY/jZGMZCZucg4uFxebzaWIGr8wM0A1UzMemuZlLRfYlYuFCo5n6qRI2NFZhwFLfuFjVdmb8",
	"1DI6Z3yM7f2B8Q9D0N/Mf5FBoGf60JrRaXdhHjzThw50zMCdM5GuJpdGZPkVrcKZPrBkOR1Tx8M5RWHz",
	"HH6ekZvklPEIH08pY2xWjFox6pozKtGti2fUwEaXmM8Um+XPBp6/dCs9i+AsOLzMrPmpow/QkvT1Mzi7",
	"9qbkL1sL9M6zne58NbRKuJwK/ov7eUmvKZlic61UYU5RDi9hLZnvx/eiEr0qrL78l53VG/JnkqoXw4p6",
	"y74je3VFy0u5Oqtj9lvauRCGP6ELr+cgVuHzK9PmDf9pszJun12iBz7Nfk7XB2MckKFQXsAeLWLAV0yt",
	"+/4yCg7PIYH3JiNHN7HH9hP8Y+wy2krdhEQ3dAt69lM33mGI+9aBYzhx8abR94/XRq3eYTREfgLvTQ0r",
	"3CPu9eF/FvcvdDMFyRn676MpxEHbriKZ2BuAtSb89H8TsuB59UoolYFObNhAXLJHBacPfAuFwUXb2LQR",
	"xA+L1K+kk3cn/icAAP//VxIHGk0hAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
