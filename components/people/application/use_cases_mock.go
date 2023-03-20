// Code generated by mockery v2.16.0. DO NOT EDIT.

package application

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// UseCasesMock is an autogenerated mock type for the UseCasesMock type
type UseCasesMock struct {
	mock.Mock
}

// GetByName provides a mock function with given fields: ctx, name
func (_m *UseCasesMock) GetByName(ctx context.Context, name string) (*PersonDTO, error) {
	ret := _m.Called(ctx, name)

	var r0 *PersonDTO
	if rf, ok := ret.Get(0).(func(context.Context, string) *PersonDTO); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*PersonDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUseCases interface {
	mock.TestingT
	Cleanup(func())
}

// NewUseCases creates a new instance of UseCasesMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUseCases(t mockConstructorTestingTNewUseCases) *UseCasesMock {
	mock := &UseCasesMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}