package server

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/hinha/PAM-Trello/app/trello"
)

type apiDashboardHandler struct {
	s trello.Service

	logger *log.Entry
}

func (h *apiDashboardHandler) performance(ctx echo.Context) error {

	verify := ctx.Get("verify")
	if verify == nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "bad payload"})
	}
	claim := verify.(map[string]interface{})

	performance, err := h.s.Performance(claim["id"].(string))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, performance)
}

func (h *apiDashboardHandler) verify(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		credential := ctx.QueryParam("key")

		data, err := h.s.Authorize(credential)
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{"error": err.Error()})
		}

		ctx.Set("verify", data)

		return next(ctx)
	}
}
