package courier

import (
	"errors"
	"time"

	"github.com/zakiafada32/retail/business/courier"
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

func (pr *CourierRepository) CreateNewCourierProvider(provider courier.CourierProvider) error {
	if err := pr.db.Where("name = ?", provider.Name).First(&CourierProvider{}).Error; err == nil {
		return errors.New("the courier provider name already exist")
	}

	providerData := convertToCourierProviderModel(provider)

	if err := pr.db.Create(&providerData).Error; err != nil {
		return err
	}

	return nil
}

func convertToCourierProviderModel(provider courier.CourierProvider) CourierProvider {
	return CourierProvider{
		ID:          provider.ID,
		Name:        provider.Name,
		Description: provider.Description,
		CreatedAt:   provider.CreatedAt,
		UpdatedAt:   provider.UpdatedAt,
	}
}
