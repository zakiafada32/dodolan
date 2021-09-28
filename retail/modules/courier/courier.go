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

func (repo *CourierRepository) FindAll() ([]courierBusiness.CourierProvider, error) {
	var couriers []CourierProvider
	err := repo.db.Find(&couriers).Error
	if err != nil {
		return []courierBusiness.CourierProvider{}, err
	}

	couriersData := make([]courierBusiness.CourierProvider, len(couriers))
	for i, courier := range couriers {
		couriersData[i] = convertToCourierProviderBusiness(courier)
	}
	return couriersData, nil
}

func (repo *CourierRepository) CreateNew(provider courierBusiness.CourierProvider) error {
	err := repo.db.Where("name = ?", provider.Name).First(&CourierProvider{}).Error
	if err == nil {
		return errors.New("the courier provider name already exist")
	}

	providerData := convertToCourierProviderModel(provider)

	err = repo.db.Create(&providerData).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *CourierRepository) Update(id uint32, name, description string) (courierBusiness.CourierProvider, error) {
	var courier CourierProvider
	err := repo.db.Where("id = ?", id).First(&courier).Error
	if err != nil {
		return courierBusiness.CourierProvider{}, err
	}

	if len(name) > 0 && name != courier.Name {
		err = repo.db.Where("name = ?", name).First(&courier).Error
		if err == nil {
			return courierBusiness.CourierProvider{}, errors.New("the courier name already exist")
		}
	}

	err = repo.db.Model(&courier).Updates(&CourierProvider{Name: name, Description: description}).Error
	if err != nil {
		return courierBusiness.CourierProvider{}, err
	}

	courierData := convertToCourierProviderBusiness(courier)
	return courierData, nil
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

func convertToCourierProviderBusiness(provider CourierProvider) courierBusiness.CourierProvider {
	return courierBusiness.CourierProvider{
		ID:          provider.ID,
		Name:        provider.Name,
		Description: provider.Description,
		CreatedAt:   provider.CreatedAt,
		UpdatedAt:   provider.UpdatedAt,
	}
}
