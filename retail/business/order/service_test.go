package order_test

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zakiafada32/retail/business"
	"github.com/zakiafada32/retail/business/order"
	"github.com/zakiafada32/retail/business/order/mocks"
)

const (
	orderId     uint32 = 1
	userId      string = "1"
	totalAmount uint64 = 1000000
)

var (
	orderService    order.Service
	orderRepository mocks.Repository
	orderData       order.Order
	ordersData      []order.Order
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFindAll(t *testing.T) {
	t.Run("Expect found all order", func(t *testing.T) {
		orderRepository.On("FindAll", userId).Return(ordersData, nil).Once()
		orders, err := orderService.FindAll(userId)
		assert.Nil(t, err)
		assert.IsType(t, []order.Order{}, orders)
	})

	t.Run("Expect internal server error when cannot fetch orders from database", func(t *testing.T) {
		orderRepository.On("FindAll", userId).Return([]order.Order{}, errors.New("error")).Once()
		_, err := orderService.FindAll(userId)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), business.NotFound)
	})
}

func TestFindById(t *testing.T) {
	t.Run("Expect found the order by id", func(t *testing.T) {
		orderRepository.On("FindById", userId, orderId).Return(orderData, nil).Once()
		orders, err := orderService.FindById(userId, orderId)
		assert.Nil(t, err)
		assert.IsType(t, order.Order{}, orders)
	})

	t.Run("Expect not found when cannot find the orders id", func(t *testing.T) {
		orderRepository.On("FindById", userId, orderId).Return(order.Order{}, errors.New("error")).Once()
		_, err := orderService.FindById(userId, orderId)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), business.NotFound)
	})
}

func TestPayment(t *testing.T) {
	t.Run("Expect payment success", func(t *testing.T) {
		orderRepository.On("Payment", userId, orderId, totalAmount).Return(nil).Once()
		err := orderService.Payment(userId, orderId, totalAmount)
		assert.Nil(t, err)
	})

	t.Run("Expect err when cannot find the orders id", func(t *testing.T) {
		orderRepository.On("Payment", userId, orderId, totalAmount).Return(errors.New("error")).Once()
		err := orderService.Payment(userId, orderId, totalAmount)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), business.BadRequest)
	})
}

func TestCourier(t *testing.T) {
	t.Run("Expect courier success", func(t *testing.T) {
		orderRepository.On("Courier", userId, orderId).Return(nil).Once()
		err := orderService.Courier(userId, orderId)
		assert.Nil(t, err)
	})

	t.Run("Expect err when cannot find the orders id", func(t *testing.T) {
		orderRepository.On("Courier", userId, orderId).Return(errors.New("error")).Once()
		err := orderService.Courier(userId, orderId)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), business.BadRequest)
	})
}

func setup() {
	orderData = order.Order{
		ID:          orderId,
		TotalAmount: totalAmount,
	}

	ordersData = []order.Order{
		{
			ID:          1,
			TotalAmount: totalAmount,
		},
		{
			ID:          2,
			TotalAmount: totalAmount,
		},
	}
	orderService = order.NewOrderService(&orderRepository)
}
