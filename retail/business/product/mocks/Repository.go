// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	product "github.com/zakiafada32/retail/business/product"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateNew provides a mock function with given fields: _a0
func (_m *Repository) CreateNew(_a0 product.Product) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(product.Product) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields:
func (_m *Repository) FindAll() ([]product.ProductAtt, error) {
	ret := _m.Called()

	var r0 []product.ProductAtt
	if rf, ok := ret.Get(0).(func() []product.ProductAtt); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.ProductAtt)
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

// FindByCategory provides a mock function with given fields: categoryId
func (_m *Repository) FindByCategory(categoryId uint32) ([]product.ProductAtt, error) {
	ret := _m.Called(categoryId)

	var r0 []product.ProductAtt
	if rf, ok := ret.Get(0).(func(uint32) []product.ProductAtt); ok {
		r0 = rf(categoryId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.ProductAtt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint32) error); ok {
		r1 = rf(categoryId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: id
func (_m *Repository) FindById(id uint32) (product.ProductAtt, error) {
	ret := _m.Called(id)

	var r0 product.ProductAtt
	if rf, ok := ret.Get(0).(func(uint32) product.ProductAtt); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(product.ProductAtt)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint32) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, updateData
func (_m *Repository) Update(id uint32, updateData product.Product) (product.ProductAtt, error) {
	ret := _m.Called(id, updateData)

	var r0 product.ProductAtt
	if rf, ok := ret.Get(0).(func(uint32, product.Product) product.ProductAtt); ok {
		r0 = rf(id, updateData)
	} else {
		r0 = ret.Get(0).(product.ProductAtt)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint32, product.Product) error); ok {
		r1 = rf(id, updateData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
