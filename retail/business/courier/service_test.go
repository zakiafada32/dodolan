package courier_test

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zakiafada32/retail/business"
	"github.com/zakiafada32/retail/business/courier"
	"github.com/zakiafada32/retail/business/courier/mocks"
)

const (
	id          uint32 = 1
	name        string = "courier"
	description string = "description"
)

var (
	courierService    courier.Service
	courierRepository mocks.Repository
	courierData       courier.CourierProvider
	couriersData      []courier.CourierProvider
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFindAll(t *testing.T) {
	t.Run("Expect found all courier provider", func(t *testing.T) {
		courierRepository.On("FindAll").Return(couriersData, nil).Once()
		couriers, err := courierService.FindAll()
		assert.Nil(t, err)
		assert.IsType(t, []courier.CourierProvider{}, couriers)
	})

	t.Run("Expect internal server error when cannot fetch courier provider from database", func(t *testing.T) {
		courierRepository.On("FindAll").Return([]courier.CourierProvider{}, errors.New(business.InternalServerError)).Once()
		_, err := courierService.FindAll()
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), business.InternalServerError)
	})
}

func TestCreateNew(t *testing.T) {
	t.Run("Expect create new courier provider", func(t *testing.T) {
		courierRepository.On("CreateNew", courierData).Return(nil).Once()
		err := courierService.CreateNew(courierData)
		assert.Nil(t, err)
	})

	t.Run("Expect bad request error when courier name already exist", func(t *testing.T) {
		courierRepository.On("CreateNew", courierData).Return(errors.New(business.BadRequest)).Once()
		err := courierService.CreateNew(courierData)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), business.BadRequest)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Expect update the courier provider by the id", func(t *testing.T) {
		courierRepository.On("Update", id, name, description).Return(courierData, nil).Once()
		courier, err := courierService.Update(id, name, description)
		assert.Nil(t, err)
		assert.Equal(t, name, courier.Name)
		assert.Equal(t, description, courier.Description)
	})

	t.Run("Expect not found when cannot find courier provider id", func(t *testing.T) {
		courierRepository.On("Update", id, name, description).Return(courier.CourierProvider{}, errors.New(business.BadRequest)).Once()
		_, err := courierService.Update(id, name, description)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), business.BadRequest)
	})
}

func setup() {
	courierData = courier.CourierProvider{
		ID:          id,
		Name:        name,
		Description: description,
	}

	couriersData = []courier.CourierProvider{
		{
			ID:          1,
			Name:        "courier 1",
			Description: "description 1",
		},
		{
			ID:          2,
			Name:        "courier 2",
			Description: "description 2",
		},
	}

	courierService = courier.NewCourierService(&courierRepository)
}
