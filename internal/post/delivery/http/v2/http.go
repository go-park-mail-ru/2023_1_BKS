package v2

import (
	"context"
	"fmt"
	"net/http"
	"post/domain"
	app "post/usecase"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=../../../../../api/openapi/post/models.cfg.yml ../../../../../api/openapi/post/post.yml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=../../../../../api/openapi/post/server.cfg.yml ../../../../../api/openapi/post/post.yml

func sendUserError(ctx echo.Context, code int, message string) error {
	postErr := ErrorHTTP{
		Code:    int32(code),
		Message: message,
	}
	err := ctx.JSON(code, postErr)
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
		newUser.SecondName, newUser.FirstName, newUser.Patronimic, newUser.Password,
		newUser.PasswordCheck, newUser.Avatar)
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	return nil
}

func (a *HttpServer) UpdateUser(ctx echo.Context) error {
	var updateUser CreateUser
	err := ctx.Bind(&updateUser)
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, "Неправильный формат запроса")
	}
	err = a.command.UpdateUser.Handle(context.Background(), uuid.New(), updateUser.Email, updateUser.Login, updateUser.PhoneNumber,
		updateUser.SecondName, updateUser.FirstName, updateUser.Patronimic, updateUser.Password,
		updateUser.PasswordCheck, updateUser.Avatar) // // Значения uuid из авторизации
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
	var post domain.User
	post, err := a.query.GetUser.Handle(context.Background(), uuid.New()) // Тут должен быть uuid из авторизации
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	result := GetUser{
		ID:          post.Id.String(),
		FirstName:   post.FullName.FirstName(),
		SecondName:  post.FullName.SecondName(),
		Patronimic:  post.FullName.Patronimic(),
		PhoneNumber: post.PhoneNumber.String(),
		Avatar:      post.Avatar.String(),
	}
	return ctx.JSON(http.StatusOK, result)
}

func (a *HttpServer) FindUserByID(ctx echo.Context, id string) error {
	var post domain.User
	uuid, err := uuid.Parse(id)
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	post, err = a.query.GetUser.Handle(context.Background(), uuid)
	if err != nil {
		return sendUserError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	result := GetUser{
		ID:          post.Id.String(),
		FirstName:   post.FullName.FirstName(),
		SecondName:  post.FullName.SecondName(),
		Patronimic:  post.FullName.Patronimic(),
		PhoneNumber: post.PhoneNumber.String(),
		Avatar:      post.Avatar.String(),
	}
	return ctx.JSON(http.StatusOK, result)
}
