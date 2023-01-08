// Code generated by mockery v2.16.0. DO NOT EDIT.

package mock

import (
	chi "github.com/go-chi/chi/v5"

	mock "github.com/stretchr/testify/mock"
)

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

type Handler_Expecter struct {
	mock *mock.Mock
}

func (_m *Handler) EXPECT() *Handler_Expecter {
	return &Handler_Expecter{mock: &_m.Mock}
}

// Routes provides a mock function with given fields: router
func (_m *Handler) Routes(router chi.Router) {
	_m.Called(router)
}

// Handler_Routes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Routes'
type Handler_Routes_Call struct {
	*mock.Call
}

// Routes is a helper method to define mock.On call
//   - router chi.Router
func (_e *Handler_Expecter) Routes(router interface{}) *Handler_Routes_Call {
	return &Handler_Routes_Call{Call: _e.mock.On("Routes", router)}
}

func (_c *Handler_Routes_Call) Run(run func(router chi.Router)) *Handler_Routes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(chi.Router))
	})
	return _c
}

func (_c *Handler_Routes_Call) Return() *Handler_Routes_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTNewHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewHandler creates a new instance of Handler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHandler(t mockConstructorTestingTNewHandler) *Handler {
	mock := &Handler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
