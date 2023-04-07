package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type MealQueryType = string

const (
	MealQueryByInterval MealQueryType = "MEAL_QUERY_BY_INTERVAL"
)

type GetUserMealReq struct {
	QueryType MealQueryType `json:"query_type"`
	From      uint64        `json:"from"`
	To        uint64        `json:"to"`
}

type GetUserMealResp struct {
}

func GetUserMeal(c echo.Context) error {
	getMealReq := new(GetUserMealReq)
	if err := c.Bind(getMealReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(getMealReq); err != nil {
		return err
	}

	return nil
}

func GetUserBFPByInterval(c echo.Context) error {
	return nil
}
