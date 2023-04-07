package handler

import (
	"github.com/khang00/health/ent"
	"github.com/khang00/health/internal/store"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	store *store.Client
}

func NewUserHandler(store *store.Client) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

type MealQueryType = string

const (
	MealQueryByInterval MealQueryType = "MEAL_QUERY_BY_INTERVAL"
)

type GetUserMealReq struct {
	UserID    int           `json:"user_id" query:"user_id" validate:"required"`
	QueryType MealQueryType `json:"query_type" query:"query_type" validate:"required"`
	From      int64         `json:"from" query:"from" validate:"required"`
	To        int64         `json:"to" query:"to" validate:"required"`
}

type GetUserMealResp struct {
	Meals []*ent.Meals `json:"meals"`
}

func (h *UserHandler) GetUserMeal(c echo.Context) error {
	getMealReq := GetUserMealReq{}
	if err := c.Bind(&getMealReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&getMealReq); err != nil {
		return err
	}

	reqCtx := c.Request().Context()
	meals, err := h.store.GetMealByInterval(reqCtx, getMealReq.UserID, getMealReq.From, getMealReq.To)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, meals)
}

func (h *UserHandler) GetUserBFPByInterval(c echo.Context) error {
	return nil
}
