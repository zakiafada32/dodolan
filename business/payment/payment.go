package payment

import "time"

type PaymentProvider struct {
	ID          uint32
	Name        string `validate:"required"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
