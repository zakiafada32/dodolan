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

func (repo *PaymentRepository) CreateNewPaymentProvider(provider paymentBusiness.PaymentProvider) error {
	if err := repo.db.Where("name = ?", provider.Name).First(&PaymentProvider{}).Error; err == nil {
		return errors.New("the payment provider name already exist")
	}

	providerData := convertToPaymentProviderModel(provider)

	if err := repo.db.Create(&providerData).Error; err != nil {
		return err
	}

	return nil
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
