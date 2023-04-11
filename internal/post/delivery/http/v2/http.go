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

func sendPostError(ctx echo.Context, code int, message string) error {
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

func (a *HttpServer) CreatePost(ctx echo.Context) error {
	var newPost CreatePost
	err := ctx.Bind(&newPost)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, "Неправильный формат запроса")
	}
	err = a.command.CreatePost.Handle(context.Background(), newPost.Title, newPost.Desciption, newPost.Images, newPost.Tags)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	return nil
}

func (a *HttpServer) UpdatePost(ctx echo.Context) error {
	var updatePost CreatePost
	err := ctx.Bind(&updatePost)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, "Неправильный формат запроса")
	}
	err = a.command.UpdatePost.Handle(context.Background(), uuid.New(), updatePost.Title, updatePost.Desciption, updatePost.Images, updatePost.Tags)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	return nil
}

func (a *HttpServer) DeletePost(ctx echo.Context) error {
	err := a.command.DeletePost.Handle(context.Background(), uuid.New())
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	return nil
}

func (a HttpServer) GetPost(ctx echo.Context) error {
	var post domain.Post
	post, err := a.query.GetPost.Handle(context.Background(), uuid.New()) // Тут должен быть uuid из авторизации
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	result := GetPost{
		ID:         post.Id.String(),
		Title:      post.Title.String(),
		Desciption: post.Desciption.String(),
		Images:     post.Images.String(),
		Tags:       post.Tags.String(),
	}
	return ctx.JSON(http.StatusOK, result)
}

func (a *HttpServer) FindPostByID(ctx echo.Context, id string) error {
	var post domain.Post
	uuid, err := uuid.Parse(id)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	post, err = a.query.GetPost.Handle(context.Background(), uuid)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}
	result := GetPost{
		ID:         post.Id.String(),
		Title:      post.Title.String(),
		Desciption: post.Desciption.String(),
		Images:     post.Images.String(),
		Tags:       post.Tags.String(),
	}
	return ctx.JSON(http.StatusOK, result)
}
