package user

import (
	"github.com/khang00/health/ent"
	"github.com/labstack/echo/v4"
	"net/http"
)

type GetUserMealReq struct {
	UserID int   `json:"user_id" query:"user_id" validate:"required"`
	From   int64 `json:"from" query:"from" validate:"required"`
	To     int64 `json:"to" query:"to" validate:"required"`
}

type GetUserMealResp struct {
	Meals []*ent.Meal `json:"meals"`
}

func (h *handler) GetUserMeal(c echo.Context) error {
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

	return c.JSON(http.StatusOK, GetUserMealResp{
		Meals: meals,
	})
}
