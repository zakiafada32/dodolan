package payment_test

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zakiafada32/retail/business"
	"github.com/zakiafada32/retail/business/payment"
	"github.com/zakiafada32/retail/business/payment/mocks"
)

const (
	id          = uint32(1)
	name        = "payment"
	description = "description"
)

var (
	paymentService    payment.Service
	paymentRepository mocks.Repository
	paymentData       payment.PaymentProvider
	paymentsData      []payment.PaymentProvider
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFindAll(t *testing.T) {
	t.Run("Expect found all payment provider", func(t *testing.T) {
		paymentRepository.On("FindAll").Return(paymentsData, nil).Once()
		payments, err := paymentService.FindAll()
		assert.Nil(t, err)
		assert.IsType(t, []payment.PaymentProvider{}, payments)
	})

	t.Run("Expect internal server error when cannot fetch payment provider from database", func(t *testing.T) {
		paymentRepository.On("FindAll").Return([]payment.PaymentProvider{}, errors.New(business.InternalServerError)).Once()
		_, err := paymentService.FindAll()
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), business.InternalServerError)
	})
}

func TestCreateNew(t *testing.T) {
	t.Run("Expect create new payment provider", func(t *testing.T) {
		paymentRepository.On("CreateNew", paymentData).Return(nil).Once()
		err := paymentService.CreateNew(paymentData)
		assert.Nil(t, err)
	})

	t.Run("Expect bad request error when payment name already exist", func(t *testing.T) {
		paymentRepository.On("CreateNew", paymentData).Return(errors.New(business.BadRequest)).Once()
		err := paymentService.CreateNew(paymentData)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), business.BadRequest)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Expect update the payment provider by the id", func(t *testing.T) {
		paymentRepository.On("Update", id, name, description).Return(paymentData, nil).Once()
		payment, err := paymentService.Update(id, name, description)
		assert.Nil(t, err)
		assert.Equal(t, name, payment.Name)
		assert.Equal(t, description, payment.Description)
	})

	t.Run("Expect not found when cannot find payment provider id", func(t *testing.T) {
		paymentRepository.On("Update", id, name, description).Return(payment.PaymentProvider{}, errors.New(business.BadRequest)).Once()
		_, err := paymentService.Update(id, name, description)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), business.BadRequest)
	})
}

func setup() {
	paymentData = payment.PaymentProvider{
		ID:          id,
		Name:        name,
		Description: description,
	}

	paymentsData = []payment.PaymentProvider{
		{
			ID:          1,
			Name:        "payment 1",
			Description: "description 1",
		},
		{
			ID:          2,
			Name:        "payment 2",
			Description: "description 2",
		},
	}

	paymentService = payment.NewPaymentService(&paymentRepository)
}
