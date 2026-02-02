package database

import (
	"context"

	"go-microservices/internal/models"
)

func (c Client) GetAllProducts(ctx context.Context, name string) ([]models.Product, error) {
	var products []models.Product
	result := c.DB.WithContext(ctx).
		Where(models.Product{Name: name}).
		Find(&products)
	return products, result.Error
}
