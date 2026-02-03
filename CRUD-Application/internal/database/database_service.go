package database

import (
	"context"
	"errors"

	"go-microservices/internal/dberrors"
	"go-microservices/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (c Client) GetAllServices(ctx context.Context, name string) ([]models.Service, error) {
	var services []models.Service
	result := c.DB.WithContext(ctx).
		Where(models.Service{Name: name}).
		Find(&services)
	return services, result.Error
}

func (c Client) AddService(ctx context.Context, service *models.Service) (*models.Service, error) {
	service.ServiceID = uuid.NewString()
	result := c.DB.WithContext(ctx).Create(&service)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	return service, nil
}

func (c Client) GetServiceByID(ctx context.Context, serviceID string) (*models.Service, error) {
	service := &models.Service{}
	result := c.DB.WithContext(ctx).
		Where(&models.Service{ServiceID: serviceID}).
		First(service)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "service", ID: serviceID}
		}
		return nil, result.Error
	}
	return service, nil
}

func (c Client) UpdateService(ctx context.Context, service *models.Service) (*models.Service, error) {
	result := c.DB.WithContext(ctx).
		Save(service)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, &dberrors.NotFoundError{Entity: "service", ID: service.ServiceID}
	}
	return service, nil
}

func (c Client) DeleteService(ctx context.Context, serviceID string) error {
	result := c.DB.WithContext(ctx).
		Delete(&models.Service{ServiceID: serviceID})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return &dberrors.NotFoundError{Entity: "service", ID: serviceID}
	}
	return nil
}
