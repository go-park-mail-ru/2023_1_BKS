package controller

import (
	"net/http"

	"github.com/2023_1_BKS/internal/app/service"
)

type Handler struct {
	userUseCase service.IUserUseCase
}

func NewHandler(useCase service.IUserUseCase) *Handler {
	return &Handler{
		userUseCase: useCase,
	}
}

func HTTPEndPoints(router *http.ServeMux, uuc service.IUserUseCase) {
	h := NewHandler(uuc)

	router.HandleFunc("/appuniq/register", h.Register)
	router.HandleFunc("/appuniq/login", h.Login)
	router.HandleFunc("/appuniq/profile", h.GetCurrentUser)
	router.HandleFunc("/appuniq/logout", h.Logout)
	router.HandleFunc("/appuniq/post", h.GetCurrentPost)
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) GetCurrentPost(w http.ResponseWriter, r *http.Request) {
}
