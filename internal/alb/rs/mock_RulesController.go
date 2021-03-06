// Code generated by mockery v1.0.0. DO NOT EDIT.

package rs

import context "context"
import elbv2 "github.com/aws/aws-sdk-go/service/elbv2"
import mock "github.com/stretchr/testify/mock"

// MockRulesController is an autogenerated mock type for the RulesController type
type MockRulesController struct {
	mock.Mock
}

// Reconcile provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockRulesController) Reconcile(_a0 context.Context, _a1 *Rules, _a2 *elbv2.Listener) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *Rules, *elbv2.Listener) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
