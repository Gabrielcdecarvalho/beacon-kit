// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Node is an autogenerated mock type for the Node type
type Node[ContextT interface{}] struct {
	mock.Mock
}

type Node_Expecter[ContextT interface{}] struct {
	mock *mock.Mock
}

func (_m *Node[ContextT]) EXPECT() *Node_Expecter[ContextT] {
	return &Node_Expecter[ContextT]{mock: &_m.Mock}
}

// CreateQueryContext provides a mock function with given fields: height, prove
func (_m *Node[ContextT]) CreateQueryContext(height int64, prove bool) (ContextT, error) {
	ret := _m.Called(height, prove)

	if len(ret) == 0 {
		panic("no return value specified for CreateQueryContext")
	}

	var r0 ContextT
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, bool) (ContextT, error)); ok {
		return rf(height, prove)
	}
	if rf, ok := ret.Get(0).(func(int64, bool) ContextT); ok {
		r0 = rf(height, prove)
	} else {
		r0 = ret.Get(0).(ContextT)
	}

	if rf, ok := ret.Get(1).(func(int64, bool) error); ok {
		r1 = rf(height, prove)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Node_CreateQueryContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateQueryContext'
type Node_CreateQueryContext_Call[ContextT interface{}] struct {
	*mock.Call
}

// CreateQueryContext is a helper method to define mock.On call
//   - height int64
//   - prove bool
func (_e *Node_Expecter[ContextT]) CreateQueryContext(height interface{}, prove interface{}) *Node_CreateQueryContext_Call[ContextT] {
	return &Node_CreateQueryContext_Call[ContextT]{Call: _e.mock.On("CreateQueryContext", height, prove)}
}

func (_c *Node_CreateQueryContext_Call[ContextT]) Run(run func(height int64, prove bool)) *Node_CreateQueryContext_Call[ContextT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(bool))
	})
	return _c
}

func (_c *Node_CreateQueryContext_Call[ContextT]) Return(_a0 ContextT, _a1 error) *Node_CreateQueryContext_Call[ContextT] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Node_CreateQueryContext_Call[ContextT]) RunAndReturn(run func(int64, bool) (ContextT, error)) *Node_CreateQueryContext_Call[ContextT] {
	_c.Call.Return(run)
	return _c
}

// NewNode creates a new instance of Node. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNode[ContextT interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *Node[ContextT] {
	mock := &Node[ContextT]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}