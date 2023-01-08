// Code generated by mockery v2.16.0. DO NOT EDIT.

package mock

import (
	chi "github.com/go-chi/chi/v5"

	mock "github.com/stretchr/testify/mock"
)

// OptionFn is an autogenerated mock type for the OptionFn type
type OptionFn struct {
	mock.Mock
}

type OptionFn_Expecter struct {
	mock *mock.Mock
}

func (_m *OptionFn) EXPECT() *OptionFn_Expecter {
	return &OptionFn_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: router
func (_m *OptionFn) Execute(router chi.Router) {
	_m.Called(router)
}

// OptionFn_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type OptionFn_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - router chi.Router
func (_e *OptionFn_Expecter) Execute(router interface{}) *OptionFn_Execute_Call {
	return &OptionFn_Execute_Call{Call: _e.mock.On("Execute", router)}
}

func (_c *OptionFn_Execute_Call) Run(run func(router chi.Router)) *OptionFn_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(chi.Router))
	})
	return _c
}

func (_c *OptionFn_Execute_Call) Return() *OptionFn_Execute_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTNewOptionFn interface {
	mock.TestingT
	Cleanup(func())
}

// NewOptionFn creates a new instance of OptionFn. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOptionFn(t mockConstructorTestingTNewOptionFn) *OptionFn {
	mock := &OptionFn{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
