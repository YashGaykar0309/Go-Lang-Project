package database

import (
	"context"

	"go-microservices/internal/models"
)

func (c Client) GetAllServices(ctx context.Context, name string) ([]models.Service, error) {
	var services []models.Service
	result := c.DB.WithContext(ctx).
		Where(models.Service{Name: name}).
		Find(&services)
	return services, result.Error
}
