// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	payment "github.com/zakiafada32/retail/business/payment"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CreateNew provides a mock function with given fields: provider
func (_m *Service) CreateNew(provider payment.PaymentProvider) error {
	ret := _m.Called(provider)

	var r0 error
	if rf, ok := ret.Get(0).(func(payment.PaymentProvider) error); ok {
		r0 = rf(provider)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields:
func (_m *Service) FindAll() ([]payment.PaymentProvider, error) {
	ret := _m.Called()

	var r0 []payment.PaymentProvider
	if rf, ok := ret.Get(0).(func() []payment.PaymentProvider); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]payment.PaymentProvider)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, name, description
func (_m *Service) Update(id uint32, name string, description string) (payment.PaymentProvider, error) {
	ret := _m.Called(id, name, description)

	var r0 payment.PaymentProvider
	if rf, ok := ret.Get(0).(func(uint32, string, string) payment.PaymentProvider); ok {
		r0 = rf(id, name, description)
	} else {
		r0 = ret.Get(0).(payment.PaymentProvider)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint32, string, string) error); ok {
		r1 = rf(id, name, description)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
