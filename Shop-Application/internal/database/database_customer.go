package database

import (
	"context"
	"errors"

	"go-microservices/internal/dberrors"
	"go-microservices/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (c Client) GetAllCustomers(ctx context.Context, emailAddress string) ([]models.Customer, error) {
	var customers []models.Customer
	result := c.DB.WithContext(ctx).
		Where(models.Customer{Email: emailAddress}).
		Find(&customers)
	return customers, result.Error
}

func (c Client) AddCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	customer.CustomerID = uuid.NewString()
	result := c.DB.WithContext(ctx).Create(&customer)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	return customer, nil
}

func (c Client) GetCustomerByID(ctx context.Context, customerID string) (*models.Customer, error) {
	customer := &models.Customer{}
	result := c.DB.WithContext(ctx).
		Where(&models.Customer{CustomerID: customerID}).
		First(customer)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "customer", ID: customerID}
		}
		return nil, result.Error
	}
	return customer, nil
}

func (c Client) UpdateCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	result := c.DB.WithContext(ctx).
		Save(customer)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, &dberrors.NotFoundError{Entity: "customer", ID: customer.CustomerID}
	}
	return customer, nil
}

func (c Client) DeleteCustomer(ctx context.Context, customerID string) error {
	result := c.DB.WithContext(ctx).
		Delete(&models.Customer{CustomerID: customerID})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return &dberrors.NotFoundError{Entity: "customer", ID: customerID}
	}
	return nil
}
