package server

import (
	"go-microservices/internal/dberrors"
	"go-microservices/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *EchoServer) GetAllProducts(ctx echo.Context) error {
	name := ctx.QueryParam("name")

	products, err := s.DB.GetAllProducts(ctx.Request().Context(), name)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, products)
}

func (s *EchoServer) AddProduct(ctx echo.Context) error {
	product := new(models.Product)
	if err := ctx.Bind(product); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	product, err := s.DB.AddProduct(ctx.Request().Context(), product)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, product)
}

func (s *EchoServer) GetProductByID(ctx echo.Context) error {
	productID := ctx.Param("productID")
	product, err := s.DB.GetProductByID(ctx.Request().Context(), productID)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, product)
}

func (s *EchoServer) UpdateProduct(ctx echo.Context) error {
	product := new(models.Product)
	productID := ctx.Param("productID")
	if err := ctx.Bind(product); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	if product.ProductID != productID {
		return ctx.JSON(http.StatusBadRequest, "product ID in path and body do not match")
	}
	updatedProduct, err := s.DB.UpdateProduct(ctx.Request().Context(), product)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, updatedProduct)
}

func (s *EchoServer) DeleteProduct(ctx echo.Context) error {
	productID := ctx.Param("productID")
	err := s.DB.DeleteProduct(ctx.Request().Context(), productID)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.NoContent(http.StatusNoContent)
}
