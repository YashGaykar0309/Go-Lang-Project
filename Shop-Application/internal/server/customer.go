package server

import (
	"go-microservices/internal/dberrors"
	"go-microservices/internal/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *EchoServer) GetAllCustomers(ctx echo.Context) error {
	emailAddress := ctx.QueryParam("emailAddress")

	customers, err := s.DB.GetAllCustomers(ctx.Request().Context(), emailAddress)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, customers)
}

func (s *EchoServer) AddCustomer(ctx echo.Context) error {
	customer := new(models.Customer)
	if err := ctx.Bind(customer); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	customer, err := s.DB.AddCustomer(ctx.Request().Context(), customer)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, customer)
}

func (s *EchoServer) GetCustomerByID(ctx echo.Context) error {
	customerID := ctx.Param("customerID")

	customer, err := s.DB.GetCustomerByID(ctx.Request().Context(), customerID)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, customer)
}

func (s *EchoServer) UpdateCustomer(ctx echo.Context) error {
	customer := new(models.Customer)
	customerID := ctx.Param("customerID")
	if err := ctx.Bind(customer); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	if customer.CustomerID != customerID {
		return ctx.JSON(http.StatusBadRequest, "customer ID in path and body do not match")
	}
	updatedCustomer, err := s.DB.UpdateCustomer(ctx.Request().Context(), customer)
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
	return ctx.JSON(http.StatusOK, updatedCustomer)
}

func (s *EchoServer) DeleteCustomer(ctx echo.Context) error {
	customerID := ctx.Param("customerID")
	err := s.DB.DeleteCustomer(ctx.Request().Context(), customerID)
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
