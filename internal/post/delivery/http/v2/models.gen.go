// Package v2 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package v2

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// CreatePost defines model for CreatePost.
type CreatePost struct {
	// Description Содержание объявления.
	Description string   `json:"Description"`
	PathImages  []string `json:"PathImages"`

	// Price Цена объявления.
	Price string `json:"Price"`

	// Tag Категория.
	Tag string `json:"Tag"`

	// Title Названия объявления.
	Title string `json:"Title"`
}

// ErrorHTTP defines model for ErrorHTTP.
type ErrorHTTP struct {
	// Code Error code
	Code int32 `json:"code"`

	// Message Error message
	Message string `json:"message"`
}

// MiniPost defines model for MiniPost.
type MiniPost struct {
	// Name Имя пользователя.
	Name       string   `json:"Name"`
	PathImages []string `json:"PathImages"`

	// PostId ID объявления.
	PostId string `json:"PostId"`

	// Price Цена объявления.
	Price string `json:"Price"`

	// Title Названия объявления.
	Title string `json:"Title"`

	// UserId ID пользователя.
	UserId string `json:"UserId"`

	// Views Количество просмотров.
	Views int `json:"Views"`
}

// Post defines model for Post.
type Post struct {
	// Close Закрыто ли объявление.
	Close bool `json:"Close"`

	// Description Содержание объявления.
	Description string   `json:"Description"`
	PathImages  []string `json:"PathImages"`

	// Price Цена объявления.
	Price string `json:"Price"`

	// Tag Категория.
	Tag string `json:"Tag"`

	// Title Названия объявления.
	Title string `json:"Title"`

	// UserId Id пользователя.
	UserId string `json:"UserId"`

	// Views Количество просмотров.
	Views int `json:"Views"`
}

// CreatePostJSONRequestBody defines body for CreatePost for application/json ContentType.
type CreatePostJSONRequestBody = CreatePost

// UpdatePostJSONRequestBody defines body for UpdatePost for application/json ContentType.
type UpdatePostJSONRequestBody = CreatePost
