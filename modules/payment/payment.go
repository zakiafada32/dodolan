package payment

import (
	"errors"
	"time"

	paymentBusiness "github.com/zakiafada32/retail/business/payment"
	"gorm.io/gorm"
)

type PaymentProvider struct {
	ID          uint32 `gorm:"primaryKey"`
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{
		db: db,
	}
}

func (repo *PaymentRepository) FindAll() ([]paymentBusiness.PaymentProvider, error) {
	var couriers []PaymentProvider
	err := repo.db.Find(&couriers).Error
	if err != nil {
		return []paymentBusiness.PaymentProvider{}, err
	}

	couriersData := make([]paymentBusiness.PaymentProvider, len(couriers))
	for i, courier := range couriers {
		couriersData[i] = convertToPaymentProviderBusiness(courier)
	}
	return couriersData, nil
}

func (repo *PaymentRepository) CreateNew(provider paymentBusiness.PaymentProvider) error {
	err := repo.db.Where("name = ?", provider.Name).First(&PaymentProvider{}).Error
	if err == nil {
		return errors.New("the payment provider name already exist")
	}

	providerData := convertToPaymentProviderModel(provider)

	err = repo.db.Create(&providerData).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *PaymentRepository) Update(id uint32, name, description string) (paymentBusiness.PaymentProvider, error) {
	var courier PaymentProvider
	err := repo.db.Where("name = ?", name).First(&courier).Error
	if err == nil {
		return paymentBusiness.PaymentProvider{}, errors.New("the courier name already exist")
	}

	err = repo.db.Where("id = ?", id).First(&courier).Error
	if err != nil {
		return paymentBusiness.PaymentProvider{}, err
	}

	err = repo.db.Model(&courier).Updates(&PaymentProvider{Name: name, Description: description}).Error
	if err != nil {
		return paymentBusiness.PaymentProvider{}, err
	}

	courierData := convertToPaymentProviderBusiness(courier)
	return courierData, nil
}

func convertToPaymentProviderModel(provider paymentBusiness.PaymentProvider) PaymentProvider {
	return PaymentProvider{
		ID:          provider.ID,
		Name:        provider.Name,
		Description: provider.Description,
		CreatedAt:   provider.CreatedAt,
		UpdatedAt:   provider.UpdatedAt,
	}
}

func convertToPaymentProviderBusiness(provider PaymentProvider) paymentBusiness.PaymentProvider {
	return paymentBusiness.PaymentProvider{
		ID:          provider.ID,
		Name:        provider.Name,
		Description: provider.Description,
		CreatedAt:   provider.CreatedAt,
		UpdatedAt:   provider.UpdatedAt,
	}
}
