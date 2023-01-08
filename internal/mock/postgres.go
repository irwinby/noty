// Code generated by mockery v2.16.0. DO NOT EDIT.

package mock

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	sql "database/sql"

	sqlx "github.com/jmoiron/sqlx"
)

// Postgres is an autogenerated mock type for the Postgres type
type Postgres struct {
	mock.Mock
}

type Postgres_Expecter struct {
	mock *mock.Mock
}

func (_m *Postgres) EXPECT() *Postgres_Expecter {
	return &Postgres_Expecter{mock: &_m.Mock}
}

// ExecContext provides a mock function with given fields: ctx, query, args
func (_m *Postgres) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 sql.Result
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) sql.Result); ok {
		r0 = rf(ctx, query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sql.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, query, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Postgres_ExecContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecContext'
type Postgres_ExecContext_Call struct {
	*mock.Call
}

// ExecContext is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - args ...interface{}
func (_e *Postgres_Expecter) ExecContext(ctx interface{}, query interface{}, args ...interface{}) *Postgres_ExecContext_Call {
	return &Postgres_ExecContext_Call{Call: _e.mock.On("ExecContext",
		append([]interface{}{ctx, query}, args...)...)}
}

func (_c *Postgres_ExecContext_Call) Run(run func(ctx context.Context, query string, args ...interface{})) *Postgres_ExecContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *Postgres_ExecContext_Call) Return(_a0 sql.Result, _a1 error) *Postgres_ExecContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// NamedQueryContext provides a mock function with given fields: ctx, query, arg
func (_m *Postgres) NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error) {
	ret := _m.Called(ctx, query, arg)

	var r0 *sqlx.Rows
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) *sqlx.Rows); ok {
		r0 = rf(ctx, query, arg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqlx.Rows)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, interface{}) error); ok {
		r1 = rf(ctx, query, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Postgres_NamedQueryContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NamedQueryContext'
type Postgres_NamedQueryContext_Call struct {
	*mock.Call
}

// NamedQueryContext is a helper method to define mock.On call
//   - ctx context.Context
//   - query string
//   - arg interface{}
func (_e *Postgres_Expecter) NamedQueryContext(ctx interface{}, query interface{}, arg interface{}) *Postgres_NamedQueryContext_Call {
	return &Postgres_NamedQueryContext_Call{Call: _e.mock.On("NamedQueryContext", ctx, query, arg)}
}

func (_c *Postgres_NamedQueryContext_Call) Run(run func(ctx context.Context, query string, arg interface{})) *Postgres_NamedQueryContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(interface{}))
	})
	return _c
}

func (_c *Postgres_NamedQueryContext_Call) Return(_a0 *sqlx.Rows, _a1 error) *Postgres_NamedQueryContext_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// SelectContext provides a mock function with given fields: ctx, dest, query, args
func (_m *Postgres) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx, dest, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, string, ...interface{}) error); ok {
		r0 = rf(ctx, dest, query, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Postgres_SelectContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SelectContext'
type Postgres_SelectContext_Call struct {
	*mock.Call
}

// SelectContext is a helper method to define mock.On call
//   - ctx context.Context
//   - dest interface{}
//   - query string
//   - args ...interface{}
func (_e *Postgres_Expecter) SelectContext(ctx interface{}, dest interface{}, query interface{}, args ...interface{}) *Postgres_SelectContext_Call {
	return &Postgres_SelectContext_Call{Call: _e.mock.On("SelectContext",
		append([]interface{}{ctx, dest, query}, args...)...)}
}

func (_c *Postgres_SelectContext_Call) Run(run func(ctx context.Context, dest interface{}, query string, args ...interface{})) *Postgres_SelectContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(context.Context), args[1].(interface{}), args[2].(string), variadicArgs...)
	})
	return _c
}

func (_c *Postgres_SelectContext_Call) Return(_a0 error) *Postgres_SelectContext_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewPostgres interface {
	mock.TestingT
	Cleanup(func())
}

// NewPostgres creates a new instance of Postgres. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPostgres(t mockConstructorTestingTNewPostgres) *Postgres {
	mock := &Postgres{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
