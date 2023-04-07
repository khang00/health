package user

import "github.com/khang00/health/internal/store"

type Handler struct {
	store *store.Client
}

func NewUserHandler(store *store.Client) *Handler {
	return &Handler{
		store: store,
	}
}
