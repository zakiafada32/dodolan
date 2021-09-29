// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	order "github.com/zakiafada32/retail/business/order"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Courier provides a mock function with given fields: userId, orderId
func (_m *Repository) Courier(userId string, orderId uint32) error {
	ret := _m.Called(userId, orderId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, uint32) error); ok {
		r0 = rf(userId, orderId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: userId
func (_m *Repository) FindAll(userId string) ([]order.Order, error) {
	ret := _m.Called(userId)

	var r0 []order.Order
	if rf, ok := ret.Get(0).(func(string) []order.Order); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]order.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: userId, orderId
func (_m *Repository) FindById(userId string, orderId uint32) (order.Order, error) {
	ret := _m.Called(userId, orderId)

	var r0 order.Order
	if rf, ok := ret.Get(0).(func(string, uint32) order.Order); ok {
		r0 = rf(userId, orderId)
	} else {
		r0 = ret.Get(0).(order.Order)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, uint32) error); ok {
		r1 = rf(userId, orderId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Payment provides a mock function with given fields: userId, orderId, totalAmount
func (_m *Repository) Payment(userId string, orderId uint32, totalAmount uint64) error {
	ret := _m.Called(userId, orderId, totalAmount)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, uint32, uint64) error); ok {
		r0 = rf(userId, orderId, totalAmount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}