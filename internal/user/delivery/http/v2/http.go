package v2

import (
	"context"
	"fmt"
	"net/http"
	"user/domain"
	app "user/usecase"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=../../../../../api/openapi/user/models.cfg.yml ../../../../../api/openapi/user/user.yml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=../../../../../api/openapi/user/server.cfg.yml ../../../../../api/openapi/user/user.yml

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

func (a *HttpServer) CreateUser(ctx echo.Context) error {
	var newUser CreateUser
	err := ctx.Bind(&newUser)
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, "Неправильный формат запроса")
	}
	err = a.command.CreateUser.Handle(context.Background(), newUser.Email, newUser.Login, newUser.PhoneNumber,
		newUser.SecondName, newUser.FirstName, *newUser.Patronimic, newUser.Password,
		newUser.PasswordCheck, *newUser.Avatar)
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	return nil
}

func (a *HttpServer) DeleteUser(ctx echo.Context) error {
	return sendUserError(ctx, http.StatusBadRequest, "Произошла неизвестная ошибка")
}

func (a *HttpServer) UpdateUser(ctx echo.Context) error {
	return sendUserError(ctx, http.StatusBadRequest, "Произошла неизвестная ошибка")
}

func (a *HttpServer) FindUserByID(ctx echo.Context, id string) error {
	return sendUserError(ctx, http.StatusBadRequest, "Произошла неизвестная ошибка")
}

func (a HttpServer) GetUser(ctx echo.Context) error {
	var user domain.User
	user, err := a.query.GetUser.Handle(context.Background(), uuid.New())
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, "Произошла неизвестная ошибка")
	}
	avatar := string(user.Avatar)
	result := GetUser{
		ID:          user.Id.String(),
		FirstName:   user.FullName[1],
		SecondName:  user.FullName[0],
		Patronimic:  &user.FullName[2],
		PhoneNumber: string(user.PhoneNumber),
		Avatar:      &avatar,
	}
	return ctx.JSON(http.StatusOK, result)
}
