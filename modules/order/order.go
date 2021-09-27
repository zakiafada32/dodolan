package order

import (
	"errors"
	"time"

	orderBusiness "github.com/zakiafada32/retail/business/order"
	"github.com/zakiafada32/retail/modules/category"
	"github.com/zakiafada32/retail/modules/courier"
	"github.com/zakiafada32/retail/modules/payment"
	"github.com/zakiafada32/retail/modules/user"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (repo *OrderRepository) FindById(orderId uint32) (orderBusiness.Order, error) {
	var order Order
	err := repo.db.Preload(clause.Associations).Where("id = ?", orderId).First(&order).Error
	if err != nil {
		return orderBusiness.Order{}, err
	}

	var items []OrderItem
	err = repo.db.Preload(clause.Associations).Where("order_id = ?", orderId).Find(&items).Error
	if err != nil {
		return orderBusiness.Order{}, err
	}

	itemsData := make([]orderBusiness.OrderItem, len(items))
	for i, item := range items {
		itemsData[i] = converToOrderItemBusiness(item)
	}

	orderData := convertToOrderBusiness(order)
	orderData.Items = itemsData
	return orderData, nil
}

func (repo *OrderRepository) FindAll() ([]orderBusiness.Order, error) {
	var orders []Order
	err := repo.db.Preload(clause.Associations).Find(&orders).Error
	if err != nil {
		return []orderBusiness.Order{}, err
	}

	ordersData := make([]orderBusiness.Order, len(orders))

	for i, order := range orders {
		var items []OrderItem
		err = repo.db.Preload(clause.Associations).Where("order_id = ?", order.ID).Find(&items).Error
		if err != nil {
			return []orderBusiness.Order{}, err
		}

		itemsData := make([]orderBusiness.OrderItem, len(items))
		for i, item := range items {
			itemsData[i] = converToOrderItemBusiness(item)
		}

		ordersData[i] = convertToOrderBusiness(order)
		ordersData[i].Items = itemsData
	}

	return ordersData, nil
}

func (repo *OrderRepository) Payment(orderId uint32, totalAmount uint64) error {
	var order Order
	err := repo.db.Where("id = ?", orderId).First(&order).Error
	if err != nil {
		return err
	}
	if *order.PaymentStatus {
		return errors.New("the order already paid")
	}
	if order.TotalAmount != totalAmount {
		return errors.New("the total amount not same")
	}

	paid := true
	err = repo.db.Model(&order).Updates(Order{PaymentStatus: &paid}).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *OrderRepository) Courier(orderId uint32) error {
	var order Order
	err := repo.db.Where("id = ?", orderId).First(&order).Error
	if err != nil {
		return err
	}
	if *order.CourierStatus {
		return errors.New("the order already delivered")
	}

	delievered := true
	err = repo.db.Model(&order).Updates(Order{CourierStatus: &delievered}).Error
	if err != nil {
		return err
	}

	return nil
}

func convertToOrderBusiness(order Order) orderBusiness.Order {
	return orderBusiness.Order{
		ID:                order.ID,
		TotalAmount:       order.TotalAmount,
		PaymentProviderID: order.PaymentProviderID,
		PaymentProvider:   order.PaymentProvider.Name,
		PaymentStatus:     *order.PaymentStatus,
		CourierProviderID: order.CourierProviderID,
		CourierProvider:   order.CourierProvider.Name,
		CourierStatus:     *order.CourierStatus,
		CreatedAt:         order.CreatedAt,
		UpdatedAt:         order.UpdatedAt,
	}
}

func converToOrderItemBusiness(item OrderItem) orderBusiness.OrderItem {
	return orderBusiness.OrderItem{
		Quantity:    item.Quantity,
		TotalAmount: item.TotalAmount,
		Product: orderBusiness.OrderProduct{
			ID:          item.OrderID,
			Name:        item.Product.Name,
			Description: item.Product.Description,
		},
	}
}
