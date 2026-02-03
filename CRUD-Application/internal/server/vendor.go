package server

import (
	"go-microservices/internal/dberrors"
	"go-microservices/internal/models"
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

func (s *EchoServer) AddVendor(ctx echo.Context) error {
	vendor := new(models.Vendor)
	if err := ctx.Bind(vendor); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	vendor, err := s.DB.AddVendor(ctx.Request().Context(), vendor)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, vendor)
}

func (s *EchoServer) GetVendorByID(ctx echo.Context) error {
	vendorID := ctx.Param("vendorID")
	vendor, err := s.DB.GetVendorByID(ctx.Request().Context(), vendorID)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, vendor)
}
