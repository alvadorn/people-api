// Code generated by mockery v2.16.0. DO NOT EDIT.

package domain

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// PeopleRepositoryMock is an autogenerated mock type for the PeopleRepositoryMock type
type PeopleRepositoryMock struct {
	mock.Mock
}

// GetByName provides a mock function with given fields: ctx, name
func (_m *PeopleRepositoryMock) GetByName(ctx context.Context, name string) (*Person, error) {
	ret := _m.Called(ctx, name)

	var r0 *Person
	if rf, ok := ret.Get(0).(func(context.Context, string) *Person); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Person)
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

type mockConstructorTestingTNewPeopleRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewPeopleRepository creates a new instance of PeopleRepositoryMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPeopleRepository(t mockConstructorTestingTNewPeopleRepository) *PeopleRepositoryMock {
	mock := &PeopleRepositoryMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
