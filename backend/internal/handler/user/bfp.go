package user

import (
	"github.com/khang00/health/backend/ent"
	"github.com/labstack/echo/v4"
	"net/http"
)

type GetUserBFPReq struct {
	UserID int   `json:"user_id" query:"user_id" validate:"required"`
	From   int64 `json:"from" query:"from" validate:"required"`
	To     int64 `json:"to" query:"to" validate:"required"`
}

type GetUserBFPResp struct {
	DataPoints []*ent.BFPDataPoint
}

func (h *handler) GetUserBFPByInterval(c echo.Context) error {
	getUserBFPReq := GetUserBFPReq{}
	if err := c.Bind(&getUserBFPReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&getUserBFPReq); err != nil {
		return err
	}

	reqCtx := c.Request().Context()
	dataPoints, err := h.store.GetBFPDataPointByInterval(reqCtx, getUserBFPReq.UserID, getUserBFPReq.From, getUserBFPReq.To)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, GetUserBFPResp{
		DataPoints: dataPoints,
	})
}
