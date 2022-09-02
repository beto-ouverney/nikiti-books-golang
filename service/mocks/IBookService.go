// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	customerror "github.com/beto-ouverney/nikiti-books/customerror"
	entity "github.com/beto-ouverney/nikiti-books/entity"

	mock "github.com/stretchr/testify/mock"
)

// IBookService is an autogenerated mock type for the IBookService type
type IBookService struct {
	mock.Mock
}

// Add provides a mock function with given fields: book
func (_m *IBookService) Add(book *entity.Book) *customerror.CustomError {
	ret := _m.Called(book)

	var r0 *customerror.CustomError
	if rf, ok := ret.Get(0).(func(*entity.Book) *customerror.CustomError); ok {
		r0 = rf(book)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*customerror.CustomError)
		}
	}

	return r0
}

// Delete provides a mock function with given fields: param
func (_m *IBookService) Delete(param string) *customerror.CustomError {
	ret := _m.Called(param)

	var r0 *customerror.CustomError
	if rf, ok := ret.Get(0).(func(string) *customerror.CustomError); ok {
		r0 = rf(param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*customerror.CustomError)
		}
	}

	return r0
}

// FindAll provides a mock function with given fields:
func (_m *IBookService) FindAll() (*[]entity.Book, *customerror.CustomError) {
	ret := _m.Called()

	var r0 *[]entity.Book
	if rf, ok := ret.Get(0).(func() *[]entity.Book); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entity.Book)
		}
	}

	var r1 *customerror.CustomError
	if rf, ok := ret.Get(1).(func() *customerror.CustomError); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*customerror.CustomError)
		}
	}

	return r0, r1
}

// FindBook provides a mock function with given fields: title
func (_m *IBookService) FindBook(title string) (*entity.Book, *customerror.CustomError) {
	ret := _m.Called(title)

	var r0 *entity.Book
	if rf, ok := ret.Get(0).(func(string) *entity.Book); ok {
		r0 = rf(title)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Book)
		}
	}

	var r1 *customerror.CustomError
	if rf, ok := ret.Get(1).(func(string) *customerror.CustomError); ok {
		r1 = rf(title)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*customerror.CustomError)
		}
	}

	return r0, r1
}

// Update provides a mock function with given fields: param, book
func (_m *IBookService) Update(param string, book *entity.Book) *customerror.CustomError {
	ret := _m.Called(param, book)

	var r0 *customerror.CustomError
	if rf, ok := ret.Get(0).(func(string, *entity.Book) *customerror.CustomError); ok {
		r0 = rf(param, book)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*customerror.CustomError)
		}
	}

	return r0
}

type mockConstructorTestingTNewIBookService interface {
	mock.TestingT
	Cleanup(func())
}

// NewIBookService creates a new instance of IBookService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIBookService(t mockConstructorTestingTNewIBookService) *IBookService {
	mock := &IBookService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
