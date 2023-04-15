package v2

import (
	app "auth/usecase"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=../../../../../api/openapi/auth/models.cfg.yml ../../../../../api/openapi/auth/auth.yml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=../../../../../api/openapi/auth/server.cfg.yml ../../../../../api/openapi/auth/auth.yml

func sendUserError(ctx echo.Context, code int, message string) error {
	userErr := ErrorHTTP{
		Code:    int32(code),
		Message: message,
	}
	err := ctx.JSON(code, userErr)
	return err
}

type HttpServer struct {
	command app.Commands
	query   app.Queries
}

func (d *HttpServer) Login(ctx echo.Context) error {
	var data SignUp

	err := ctx.Bind(&data)
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, "Неправильный формат запроса")
	}

	//massive := append([]string{}, "dw")
	result, err := d.command.CreateToken.CreateJWSWithClaims([]string{}, "appUniqFront", "auth")
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	return ctx.JSON(http.StatusOK, string(result))
}
