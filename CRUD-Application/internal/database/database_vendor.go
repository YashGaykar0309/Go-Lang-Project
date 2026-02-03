package database

import (
	"context"
	"errors"

	"go-microservices/internal/dberrors"
	"go-microservices/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (c Client) GetAllVendors(ctx context.Context, name string) ([]models.Vendor, error) {
	var vendors []models.Vendor
	result := c.DB.WithContext(ctx).
		Where(models.Vendor{Name: name}).
		Find(&vendors)
	return vendors, result.Error
}

func (c Client) AddVendor(ctx context.Context, vendor *models.Vendor) (*models.Vendor, error) {
	vendor.VendorID = uuid.NewString()
	result := c.DB.WithContext(ctx).Create(&vendor)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	return vendor, nil
}

func (c Client) GetVendorByID(ctx context.Context, vendorID string) (*models.Vendor, error) {
	vendor := &models.Vendor{}
	result := c.DB.WithContext(ctx).
		Where(&models.Vendor{VendorID: vendorID}).
		First(vendor)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "vendor", ID: vendorID}
		}
		return nil, result.Error
	}
	return vendor, nil
}
