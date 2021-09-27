package order

import (
	"time"

	"github.com/zakiafada32/retail/modules/category"
	"github.com/zakiafada32/retail/modules/courier"
	"github.com/zakiafada32/retail/modules/payment"
	"github.com/zakiafada32/retail/modules/user"
)

type Order struct {
	ID                uint32
	UserID            string
	User              user.User
	TotalAmount       uint64
	PaymentProviderID uint32
	PaymentProvider   payment.PaymentProvider
	PaymentStatus     *bool
	CourierProviderID uint32
	CourierProvider   courier.CourierProvider
	CourierStatus     *bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type OrderItem struct {
	OrderID     uint32
	Order       Order
	ProductID   uint32
	Product     category.Product
	Quantity    uint32
	TotalAmount uint64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
