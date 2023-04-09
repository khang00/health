package user

import (
	"github.com/khang00/health/backend/ent"
	"github.com/khang00/health/backend/internal/store"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler interface {
	GetUser(c echo.Context) error
	GetUserBFPByInterval(c echo.Context) error
	GetUserMeal(c echo.Context) error
}

type handler struct {
	store store.Client
}

func NewUserHandler(store store.Client) Handler {
	return &handler{
		store: store,
	}
}

type GetUserReq struct {
	UserName string `json:"user_name" query:"user_name" validate:"required"`
}

type GetUserResponse struct {
	User         *ent.User          `json:"user"`
	Achievements []*ent.Achievement `json:"achievements"`
}

func (h *handler) GetUser(c echo.Context) error {
	getUserReq := GetUserReq{}
	if err := c.Bind(&getUserReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&getUserReq); err != nil {
		return err
	}

	reqCtx := c.Request().Context()
	user, err := h.store.GetUserWithAchievement(reqCtx, getUserReq.UserName)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, GetUserResponse{
		User:         user,
		Achievements: user.Edges.Achievements,
	})
}
