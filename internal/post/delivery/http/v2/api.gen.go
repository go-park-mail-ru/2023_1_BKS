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
	// Удалить пост.
	// (DELETE /post)
	DeletePost(ctx echo.Context) error
	// Вернуть данные поста.
	// (GET /post)
	GetPost(ctx echo.Context) error
	// Создать новый пост.
	// (POST /post)
	CreatePost(ctx echo.Context) error
	// Обновить пост.
	// (PUT /user)
	UpdatePost(ctx echo.Context) error
	// Вернуть пост по id.
	// (GET /post/{id})
	FindPostByID(ctx echo.Context, id string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// DeletePost converts echo context to params.
func (w *ServerInterfaceWrapper) DeletePost(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeletePost(ctx)
	return err
}

// GetPost converts echo context to params.
func (w *ServerInterfaceWrapper) GetPost(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPost(ctx)
	return err
}

// CreatePost converts echo context to params.
func (w *ServerInterfaceWrapper) CreatePost(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreatePost(ctx)
	return err
}

// UpdatePost converts echo context to params.
func (w *ServerInterfaceWrapper) UpdatePost(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdatePost(ctx)
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

	router.DELETE(baseURL+"/post", wrapper.DeletePost)
	router.GET(baseURL+"/post", wrapper.GetPost)
	router.POST(baseURL+"/post", wrapper.CreatePost)
	router.PUT(baseURL+"/post", wrapper.UpdatePost)
	router.GET(baseURL+"/post/:id", wrapper.FindPostByID)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xX3W4bRRR+ldXA5cp2G672Lm0oWEJtJJqrkovp7tge2J2ZzswGRZalJrmgEpUCXCEh",
	"BSF4ADfUqqnr7SuceSN0Zr34b5NglFapxN1mZnJ+vnO+7xz3SSwzJQUT1pCoT0zcYxn1n3c1o5btGabx",
	"L5qmDzoketQnSkvFtOXMv9o+oJb6FwkzsebKcilIROAHOIehO4ahexrAWyhg4p7DKyhmxyOYuNMGCYk9",
	"VIxExFjNRZcMQvJpRnm6btAfb2bpHtfG3qcZqwnvZ3jjTjcz94XsclFj6hco4E8Yw3Qzc7vUmG+lTmos",
	"/oqolab+m827PRZ/U2fYmz2HkXsKr2GIxitPp5t6sloKnvG4xs2ZO3bfwcgduWM4h2JDyz0p2P08e8z0",
	"9YH9JYulSC7ohT9gCG9gDBMYbwbDICSaPcm5ZgmJHs06t2qU5UyWQliCb7FPF7pitZhhRbX9wT6yRGup",
	"P3/4cBfzWaZkLJOaLP0/BP4uJB2pM2pJRLiwW7fniXFhWZdpRCxjxtDuhYaq66swmTmsnmPonzH7/lXl",
	"mrWgvXOdQvA/lVbbpr1zOX+WKbPApQWSoE0uOnI9zAeKiW3Fg61GK4CXXvzckRfFcxi7IxgGuWEaz4bw",
	"F3ZbsK3UnuBPfODcphj57Cj4Km+1tmLD9AGPWYCd7U9YsL3bJiE5YNqUXm81Wo0WIigVE1RxEpEtfxQS",
	"RW3PN34zN1WtUmY9wsgLioG3ExKRHX/u+YOIGSWFKTlzu/VJ7Sgp1vB/HrgTeAlDmLgfYdrAkBLWoXlq",
	"S/kQlgn/SZVKeeydN782Usw3BPz6WLMOichHzfkK0ZztD825PPky1I6hMYblnsEEBxF+jOEFTqWG7waT",
	"ZxnVh/j89zJYGLvjSwfiICRdZtcxqwRnDbDWteW7sCzVJXwGr28+zD95BkzdSQkzYj6FqfseRpdjrqSp",
	"AX0BkZLazNg7Mjl8Z5DP5cPqnA3Win2rnh01aSE/3BG8hZF7BlMoUAgKeFUicvPr+FsVbFnHqc8Npb64",
	"oo55TRn3VHKzyvjeOPsvewMKeDHD+AOR07N/Ar5SUAdhOZGafZ4MMMhafb3HRYIw3jn0Y1tRTTNmmTZ+",
	"s1sOtr1z6VaAP7D8NCQhEX7FIDwhq10RLsC2ukfsv8OOqQbJJu0yy/YEt7gPoTtWpsAFtfIXAU8apUvc",
	"fqp65zolEelZq0zUbPZ70lis5KCJK09IDqjm9HFaFqe6LKV5hgpJZUxTvELj+4O/AwAA//9yQEeGpBAA",
	"AA==",
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
