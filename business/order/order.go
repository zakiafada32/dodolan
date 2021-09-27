package order

import (
	"time"
)

type Order struct {
	ID                uint32      `json:"id"`
	TotalAmount       uint64      `json:"total_amount"`
	PaymentProviderID uint32      `json:"payment_provider_id"`
	PaymentProvider   string      `json:"payment_provider"`
	PaymentStatus     bool        `json:"payment_status"`
	CourierProviderID uint32      `json:"courier_provider_id"`
	CourierProvider   string      `json:"courier_provider"`
	CourierStatus     bool        `json:"status_status"`
	Items             []OrderItem `json:"items"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
}

type OrderItem struct {
	Product     OrderProduct `json:"product"`
	Quantity    uint32       `json:"quantity"`
	TotalAmount uint64       `json:"total_amount"`
}

type OrderProduct struct {
	ID          uint32 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
