package database

import (
	"context"
	"errors"

	"go-microservices/internal/dberrors"
	"go-microservices/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (c Client) GetAllProducts(ctx context.Context, name string) ([]models.Product, error) {
	var products []models.Product
	result := c.DB.WithContext(ctx).
		Where(models.Product{Name: name}).
		Find(&products)
	return products, result.Error
}

func (c Client) AddProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	product.ProductID = uuid.NewString()
	result := c.DB.WithContext(ctx).Create(&product)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	return product, nil
}

func (c Client) GetProductByID(ctx context.Context, productID string) (*models.Product, error) {
	product := &models.Product{}
	result := c.DB.WithContext(ctx).
		Where(&models.Product{ProductID: productID}).
		First(product)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "product", ID: productID}
		}
		return nil, result.Error
	}
	return product, nil
}
