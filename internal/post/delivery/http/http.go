package v2

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2023_1_BKS/internal/post/domain"

	app "github.com/go-park-mail-ru/2023_1_BKS/internal/post/usecase"

	"github.com/go-park-mail-ru/2023_1_BKS/pkg/jwt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=../../../../api/openapi/post/models.cfg.yml ../../../../api/openapi/post/post.yml
//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=../../../../api/openapi/post/server.cfg.yml ../../../../api/openapi/post/post.yml

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

// ///////////////////////////////////////////////////////////////////////////////////////////
// Посты
// ///////////////////////////////////////////////////////////////////////////////////////////

func (a *HttpServer) CreatePost(ctx echo.Context) error {
	var post CreatePost

	err := ctx.Bind(&post)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, "Incorrect request format")
	}

	headerAuth := ctx.Request().Header.Get("Authorization")
	userId := jwt.ClaimParse(headerAuth, "id")

	uuidUser, err := uuid.Parse(userId)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	postDTO := domain.Post{
		UserID:      &uuidUser,
		Title:       &post.Title,
		Description: &post.Description,
		Price:       &post.Price,
		Category:    &post.Category,
		PathImages:  &post.PathImages,
	}

	uuidPost, wrappErr := a.command.CreatePost.Handle(context.Background(), postDTO)
	if wrappErr.Error != nil {
		return sendPostError(ctx, wrappErr.HTTPStatusCode, wrappErr.Error.Error())
	}

	return ctx.JSON(http.StatusCreated, uuidPost)
}

func (a *HttpServer) UpdatePost(ctx echo.Context, id string) error {
	var post EditPost

	err := ctx.Bind(&post)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, "Неправильный формат запроса")
	}

	headerAuth := ctx.Request().Header.Get("Authorization")
	idUser := jwt.ClaimParse(headerAuth, "id")

	uuidUser, err := uuid.Parse(idUser)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, err.Error())
	}

	uuidPost, err := uuid.Parse(id)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, err.Error())
	}

	postDTO := domain.Post{
		Id:          &uuidPost,
		UserID:      &uuidUser,
		Title:       post.Title,
		Description: post.Description,
		Price:       post.Price,
		Category:    post.Category,
		PathImages:  post.PathImages,
	}

	wrappErr := a.command.UpdatePost.Handle(context.Background(), postDTO)
	if wrappErr.Error != nil {
		return sendPostError(ctx, wrappErr.HTTPStatusCode, wrappErr.Error.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (a *HttpServer) DeletePost(ctx echo.Context, id string) error {
	uuidPost, err := uuid.Parse(id)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, err.Error())
	}

	wrapErr := a.command.DeletePost.Handle(context.Background(), uuidPost)
	if wrapErr.Error != nil {
		return sendPostError(ctx, http.StatusBadRequest, err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
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

	result := FullPost{
		Category:    *resultDTO.Category,
		Status:      *resultDTO.Status,
		Description: *resultDTO.Description,
		PathImages:  *resultDTO.PathImages,
		Price:       *resultDTO.Price,
		Title:       *resultDTO.Title,
		UserId:      resultDTO.UserID.String(),
		Views:       *resultDTO.Views,
	}

	return ctx.JSON(http.StatusCreated, result)
}

func (a HttpServer) GetMiniPost(ctx echo.Context, params GetMiniPostParams) error {
	resultDTO, err := a.query.GetSortNewPost.Handle(context.Background(), 1)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	var result []MiniPost
	for _, post := range resultDTO {
		p := MiniPost{
			PostId:     post.Id.String(),
			PathImages: *post.PathImages,
			Price:      *post.Price,
			Title:      *post.Title,
			UserId:     post.UserID.String(),
			Views:      *post.Views,
		}
		result = append(result, p)
	}

	return ctx.JSON(http.StatusOK, result)
}

// ///////////////////////////////////////////////////////////////////////////////////////////
// Корзина
// ///////////////////////////////////////////////////////////////////////////////////////////
func (a HttpServer) AddCart(ctx echo.Context, id string) error {
	postId, err := uuid.Parse(id)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	headerAuth := ctx.Request().Header.Get("Authorization")
	user := jwt.ClaimParse(headerAuth, "id")

	userId, err := uuid.Parse(user)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	err = a.command.AddCart.Handle(context.Background(), userId, postId)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	return ctx.JSON(http.StatusCreated, "Ok")
}

func (a HttpServer) RemoveCart(ctx echo.Context, id string) error {
	postId, err := uuid.Parse(id)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	headerAuth := ctx.Request().Header.Get("Authorization")
	user := jwt.ClaimParse(headerAuth, "id")

	userId, err := uuid.Parse(user)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	err = a.command.RemoveCart.Handle(context.Background(), userId, postId)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	return ctx.JSON(http.StatusCreated, "Ok")
}

func (a HttpServer) GetCart(ctx echo.Context) error {

	headerAuth := ctx.Request().Header.Get("Authorization")
	user := jwt.ClaimParse(headerAuth, "id")

	userId, err := uuid.Parse(user)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	resultUUID, err := a.query.GetCart.Handle(context.Background(), userId)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	resultDTO, err := a.query.GetByArray.Handle(context.Background(), resultUUID)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	var result []MiniPost
	for _, post := range resultDTO {
		p := MiniPost{
			PostId:     post.Id.String(),
			PathImages: *post.PathImages,
			Price:      *post.Price,
			Title:      *post.Title,
			UserId:     post.UserID.String(),
			Views:      *post.Views,
		}
		result = append(result, p)
	}

	return ctx.JSON(http.StatusOK, resultDTO)
}

// ///////////////////////////////////////////////////////////////////////////////////////////
// Избранное
// ///////////////////////////////////////////////////////////////////////////////////////////

func (a HttpServer) AddFavorite(ctx echo.Context, id string) error {
	postId, err := uuid.Parse(id)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	headerAuth := ctx.Request().Header.Get("Authorization")
	user := jwt.ClaimParse(headerAuth, "id")

	userId, err := uuid.Parse(user)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	err = a.command.AddFavorite.Handle(context.Background(), userId, postId)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	return ctx.JSON(http.StatusCreated, "Ok")
}

func (a HttpServer) RemoveFavorite(ctx echo.Context, id string) error {
	postId, err := uuid.Parse(id)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	headerAuth := ctx.Request().Header.Get("Authorization")
	user := jwt.ClaimParse(headerAuth, "id")

	userId, err := uuid.Parse(user)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	err = a.command.RemoveFavorite.Handle(context.Background(), userId, postId)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	return ctx.JSON(http.StatusCreated, "Ok")
}

func (a HttpServer) GetFavorite(ctx echo.Context) error {

	headerAuth := ctx.Request().Header.Get("Authorization")
	user := jwt.ClaimParse(headerAuth, "id")
	fmt.Println(user)
	userId, err := uuid.Parse(user)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	resultUUID, err := a.query.GetFavorite.Handle(context.Background(), userId)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	resultDTO, err := a.query.GetByArray.Handle(context.Background(), resultUUID)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	var result []MiniPost
	for _, post := range resultDTO {
		p := MiniPost{
			PostId:     post.Id.String(),
			PathImages: *post.PathImages,
			Price:      *post.Price,
			Title:      *post.Title,
			UserId:     post.UserID.String(),
			Views:      *post.Views,
		}
		result = append(result, p)
	}
	fmt.Println(result)
	return ctx.JSON(http.StatusOK, result)
}

// ///////////////////////////////////////////////////////////////////////////////////////////
// Поиск
// ///////////////////////////////////////////////////////////////////////////////////////////

func (a HttpServer) Search(ctx echo.Context, params SearchParams) error {

	resultDTOUuid, err := a.query.SearhPost.Handle(context.Background(), params.Query)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	resultDTO, err := a.query.GetByArray.Handle(context.Background(), resultDTOUuid)
	if err != nil {
		return sendPostError(ctx, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	var result []MiniPost
	for _, post := range resultDTO {
		p := MiniPost{
			PostId:     post.Id.String(),
			PathImages: *post.PathImages,
			Price:      *post.Price,
			Title:      *post.Title,
			UserId:     post.UserID.String(),
			Views:      *post.Views,
		}
		result = append(result, p)
	}
	fmt.Println(result)
	return ctx.JSON(http.StatusOK, result)
}
