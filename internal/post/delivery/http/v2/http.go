package v2

import (
	"context"
	"fmt"
	"net/http"
	"pkg/jwt"
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
	var post CreatePost

	err := ctx.Bind(&post)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, "Неправильный формат запроса")
	}

	headerAuth := ctx.Request().Header.Get("Authorization")
	id := jwt.ClaimParse(headerAuth, "id")

	uuids, err := uuid.Parse(id)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	postDTO := domain.Post{
		UserID:     uuids,
		Title:      post.Title,
		Desciption: post.Description,
		Price:      post.Price,
		Tags:       post.Tag,
		PathImages: post.PathImages,
	}

	err = a.command.CreatePost.Handle(context.Background(), postDTO)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	return ctx.JSON(http.StatusCreated, "Ok")
}

func (a *HttpServer) UpdatePost(ctx echo.Context, id string) error {
	var post CreatePost

	err := ctx.Bind(&post)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, "Неправильный формат запроса")
	}

	headerAuth := ctx.Request().Header.Get("Authorization")
	idUser := jwt.ClaimParse(headerAuth, "id")

	uuidUser, err := uuid.Parse(idUser)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	uuidPost, err := uuid.Parse(id)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	postDTO := domain.Post{
		UserID:     uuidUser,
		Title:      post.Title,
		Desciption: post.Description,
		Price:      post.Price,
		Tags:       post.Tag,
		PathImages: post.PathImages,
	}

	err = a.command.UpdatePost.Handle(context.Background(), uuidPost, postDTO)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	return ctx.JSON(http.StatusCreated, "Ok")
}

func (a *HttpServer) DeletePost(ctx echo.Context, id string) error {
	uuidPost, err := uuid.Parse(id)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	err = a.command.DeletePost.Handle(context.Background(), uuidPost)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	return ctx.JSON(http.StatusCreated, "Ok")
}

func (a *HttpServer) ClosePost(ctx echo.Context, userId string, id string) error {
	uuidPost, err := uuid.Parse(id)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	err = a.command.ClosePost.Handle(context.Background(), uuidPost)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	return ctx.JSON(http.StatusCreated, "Ok")
}

func (a HttpServer) FindPostByID(ctx echo.Context, id string) error {
	uuidPost, err := uuid.Parse(id)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	resultDTO, err := a.query.GetIdPost.Handle(context.Background(), uuidPost)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	result := Post{
		Tag:         resultDTO.Tags,
		Close:       resultDTO.Close,
		Description: resultDTO.Desciption,
		PathImages:  resultDTO.PathImages,
		Price:       resultDTO.Price,
		Title:       resultDTO.Title,
		UserId:      resultDTO.UserID.String(),
		Views:       resultDTO.Views,
	}

	return ctx.JSON(http.StatusCreated, result)
}

func (a HttpServer) FindOpenPostByUserID(ctx echo.Context, idUser string, page int) error {
	uuidUser, err := uuid.Parse(idUser)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	resultDTO, err := a.query.GetUserIdOpenPost.Handle(context.Background(), uuidUser, page)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	var result []Post
	for _, post := range resultDTO {
		p := Post{
			Close:       post.Close, // В реальности надо сделать не обязательныим полем
			Description: post.Desciption,
			PathImages:  post.PathImages,
			Price:       post.Price,
			Title:       post.Title,
			UserId:      post.UserID.String(),
			Views:       post.Views,
		}
		result = append(result, p)
	}

	return ctx.JSON(http.StatusCreated, result)
}

func (a HttpServer) FindClosePostByUserID(ctx echo.Context, idUser string, page int) error {
	uuidUser, err := uuid.Parse(idUser)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	resultDTO, err := a.query.GetUserIdClosePost.Handle(context.Background(), uuidUser, page)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	var result []Post
	for _, post := range resultDTO {
		p := Post{
			Close:       post.Close, // В реальности надо сделать не обязательныим полем
			Description: post.Desciption,
			PathImages:  post.PathImages,
			Price:       post.Price,
			Title:       post.Title,
			UserId:      post.UserID.String(),
			Views:       post.Views,
		}
		result = append(result, p)
	}

	return ctx.JSON(http.StatusCreated, result)
}

func (a HttpServer) FindPostByTag(ctx echo.Context, tag string, page int) error {
	resultDTO, err := a.query.GetTagPost.Handle(context.Background(), tag, page)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	var result []MiniPost
	for _, post := range resultDTO {
		p := MiniPost{
			PostId:     post.Id.String(),
			PathImages: post.PathImages,
			Price:      post.Price,
			Title:      post.Title,
			UserId:     post.UserID.String(),
			Views:      post.Views,
		}
		result = append(result, p)
	}

	return ctx.JSON(http.StatusCreated, result)
}

func (a HttpServer) GetAllPost(ctx echo.Context, page int) error {
	resultDTO, err := a.query.GetSortNewPost.Handle(context.Background(), page)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	var result []MiniPost
	for _, post := range resultDTO {
		p := MiniPost{
			PostId:     post.Id.String(),
			PathImages: post.PathImages,
			Price:      post.Price,
			Title:      post.Title,
			UserId:     post.UserID.String(),
			Views:      post.Views,
		}
		result = append(result, p)
	}

	return ctx.JSON(http.StatusCreated, result)
}
