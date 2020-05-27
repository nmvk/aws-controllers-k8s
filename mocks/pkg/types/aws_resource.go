// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	types "github.com/aws/aws-service-operator-k8s/pkg/types"
	mock "github.com/stretchr/testify/mock"
)

// AWSResource is an autogenerated mock type for the AWSResource type
type AWSResource struct {
	mock.Mock
}

// AccountID provides a mock function with given fields:
func (_m *AWSResource) AccountID() types.AWSAccountID {
	ret := _m.Called()

	var r0 types.AWSAccountID
	if rf, ok := ret.Get(0).(func() types.AWSAccountID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(types.AWSAccountID)
	}

	return r0
}

// IsDeleted provides a mock function with given fields:
func (_m *AWSResource) IsDeleted() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
