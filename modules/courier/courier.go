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

func NewCourierProvider(courierProvider courier.CourierProvider) *CourierProvider {
	return &CourierProvider{
		ID:          courierProvider.ID,
		Name:        courierProvider.Name,
		Description: courierProvider.Description,
		CreatedAt:   courierProvider.CreatedAt,
		UpdatedAt:   courierProvider.UpdatedAt,
	}
}

func (pr *CourierRepository) CreateNewCourierProvider(courierProvider courier.CourierProvider) error {
	if err := pr.db.Where("name = ?", courierProvider.Name).First(&CourierProvider{}).Error; err == nil {
		return errors.New("the courier provider name already exist")
	}

	courierProviderData := NewCourierProvider(courierProvider)

	if err := pr.db.Create(courierProviderData).Error; err != nil {
		return err
	}

	return nil
}
