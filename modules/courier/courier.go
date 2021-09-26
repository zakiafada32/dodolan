package courier

import (
	"errors"
	"time"

	courierBusiness "github.com/zakiafada32/retail/business/courier"
	"gorm.io/gorm"
)

type CourierProvider struct {
	ID          uint32 `gorm:"primaryKey"`
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CourierRepository struct {
	db *gorm.DB
}

func NewCourierRepository(db *gorm.DB) *CourierRepository {
	return &CourierRepository{
		db: db,
	}
}

func (repo *CourierRepository) CreateNewCourierProvider(provider courierBusiness.CourierProvider) error {
	if err := repo.db.Where("name = ?", provider.Name).First(&CourierProvider{}).Error; err == nil {
		return errors.New("the courier provider name already exist")
	}

	providerData := convertToCourierProviderModel(provider)

	if err := repo.db.Create(&providerData).Error; err != nil {
		return err
	}

	return nil
}

func convertToCourierProviderModel(provider courierBusiness.CourierProvider) CourierProvider {
	return CourierProvider{
		ID:          provider.ID,
		Name:        provider.Name,
		Description: provider.Description,
		CreatedAt:   provider.CreatedAt,
		UpdatedAt:   provider.UpdatedAt,
	}
}
