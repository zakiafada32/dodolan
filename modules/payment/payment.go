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

func NewPaymentProvider(paymentProvider payment.PaymentProvider) *PaymentProvider {
	return &PaymentProvider{
		ID:          paymentProvider.ID,
		Name:        paymentProvider.Name,
		Description: paymentProvider.Description,
		CreatedAt:   paymentProvider.CreatedAt,
		UpdatedAt:   paymentProvider.UpdatedAt,
	}
}

func (pr *PaymentRepository) CreateNewPaymentProvider(paymentProvider payment.PaymentProvider) error {
	if err := pr.db.Where("name = ?", paymentProvider.Name).First(&PaymentProvider{}).Error; err == nil {
		return errors.New("the payment provider name already exist")
	}

	paymentProviderData := NewPaymentProvider(paymentProvider)

	if err := pr.db.Create(paymentProviderData).Error; err != nil {
		return err
	}

	return nil
}
