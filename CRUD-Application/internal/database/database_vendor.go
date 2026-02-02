package database

import (
	"context"

	"go-microservices/internal/models"
)

func (c Client) GetAllVendors(ctx context.Context, name string) ([]models.Vendor, error) {
	var vendors []models.Vendor
	result := c.DB.WithContext(ctx).
		Where(models.Vendor{Name: name}).
		Find(&vendors)
	return vendors, result.Error
}
