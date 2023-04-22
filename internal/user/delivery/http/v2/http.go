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

	user := domain.User{
		Id: uuid.New(),

		Email:       newUser.Email,
		PhoneNumber: newUser.PhoneNumber,

		Login:    newUser.Login,
		Password: newUser.Password,

		Name: newUser.Name,

		PathToAvatar: newUser.Avatar,
	}

	err = a.command.CreateUser.Handle(context.Background(), newUser.PasswordCheck, user)
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	return ctx.JSON(http.StatusCreated, 23424)
}

func (a *HttpServer) UpdateUser(ctx echo.Context) error {
	var updateUser CreateUser

	err := ctx.Bind(&updateUser)
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, "Неправильный формат запроса")
	}

	user := domain.User{
		Id: uuid.New(), // Значения uuid из авторизации

		Email:       updateUser.Email,
		PhoneNumber: updateUser.PhoneNumber,

		Login:    updateUser.Login,
		Password: updateUser.Password,

		Name: updateUser.Name,

		PathToAvatar: updateUser.Avatar,
	}

	err = a.command.UpdateUser.Handle(context.Background(), user)
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	return nil
}

func (a *HttpServer) DeleteUser(ctx echo.Context) error {
	err := a.command.DeleteUser.Handle(context.Background(), uuid.New()) // Значения из авторизации
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	return nil
}

func (a HttpServer) GetUser(ctx echo.Context) error {
	var user domain.User
	user, err := a.query.GetUser.Handle(context.Background(), uuid.New()) // Тут должен быть uuid из авторизации
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	result := CreateUser{
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Avatar:      user.PathToAvatar,
	}
	return ctx.JSON(http.StatusOK, result)
}

func (a *HttpServer) FindUserByID(ctx echo.Context, id string) error {
	var user domain.User
	uuid, err := uuid.Parse(id)
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	user, err = a.query.GetUser.Handle(context.Background(), uuid)
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	result := GetUser{
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Avatar:      user.PathToAvatar,
	}
	return ctx.JSON(http.StatusOK, result)
}
