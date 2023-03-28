package post

import (
	"github.com/go-park-mail-ru/2023_1_BKS/golang-rest-api/internal/adapters"
	"github.com/go-park-mail-ru/2023_1_BKS/golang-rest-api/internal/domain/post"
	"github.com/gorilla/mux"
)

const (
	postURL  = "/posts/:post_id"
	postsURL = "/posts"
)

type handler struct {
	postService post.Service
}

func NewHandler(service post.Service) adapters.Handler {
	return &handler{postService: service}
}

func (h *handler) Register(router *mux.Router) {
	panic("not implemented")
}
