package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *EchoServer) GetAllServices(ctx echo.Context) error {
	name := ctx.QueryParam("name")

	services, err := s.DB.GetAllServices(ctx.Request().Context(), name)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, services)
}
