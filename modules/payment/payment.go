package payment

import (
	"errors"
	"time"

	"github.com/zakiafada32/retail/business/payment"
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

func (pr *PaymentRepository) CreateNewPaymentProvider(provider payment.PaymentProvider) error {
	if err := pr.db.Where("name = ?", provider.Name).First(&PaymentProvider{}).Error; err == nil {
		return errors.New("the payment provider name already exist")
	}

	providerData := convertToPaymentProviderModel(provider)

	if err := pr.db.Create(&providerData).Error; err != nil {
		return err
	}

	return nil
}

func convertToPaymentProviderModel(provider payment.PaymentProvider) PaymentProvider {
	return PaymentProvider{
		ID:          provider.ID,
		Name:        provider.Name,
		Description: provider.Description,
		CreatedAt:   provider.CreatedAt,
		UpdatedAt:   provider.UpdatedAt,
	}
}
