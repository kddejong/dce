// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import account "github.com/Optum/dce/pkg/account"
import mock "github.com/stretchr/testify/mock"

// Eventer is an autogenerated mock type for the Eventer type
type Eventer struct {
	mock.Mock
}

// AccountCreate provides a mock function with given fields: _a0
func (_m *Eventer) AccountCreate(_a0 *account.Account) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*account.Account) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AccountDelete provides a mock function with given fields: _a0
func (_m *Eventer) AccountDelete(_a0 *account.Account) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*account.Account) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AccountReset provides a mock function with given fields: _a0
func (_m *Eventer) AccountReset(_a0 *account.Account) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*account.Account) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AccountUpdate provides a mock function with given fields: _a0
func (_m *Eventer) AccountUpdate(_a0 *account.Account) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*account.Account) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
