package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *EchoServer) GetAllVendors(ctx echo.Context) error {
	name := ctx.QueryParam("name")

	vendors, err := s.DB.GetAllVendors(ctx.Request().Context(), name)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, vendors)
}
