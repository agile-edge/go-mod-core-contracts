// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/agile-edgex/go-mod-core-contracts/v3/dtos/common"

	errors "github.com/agile-edgex/go-mod-core-contracts/v3/errors"

	mock "github.com/stretchr/testify/mock"
)

// CommonClient is an autogenerated mock type for the CommonClient type
type CommonClient struct {
	mock.Mock
}

// AddSecret provides a mock function with given fields: ctx, request
func (_m *CommonClient) AddSecret(ctx context.Context, request common.SecretRequest) (common.BaseResponse, errors.EdgeX) {
	ret := _m.Called(ctx, request)

	var r0 common.BaseResponse
	if rf, ok := ret.Get(0).(func(context.Context, common.SecretRequest) common.BaseResponse); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(common.BaseResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, common.SecretRequest) errors.EdgeX); ok {
		r1 = rf(ctx, request)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// Configuration provides a mock function with given fields: ctx
func (_m *CommonClient) Configuration(ctx context.Context) (common.ConfigResponse, errors.EdgeX) {
	ret := _m.Called(ctx)

	var r0 common.ConfigResponse
	if rf, ok := ret.Get(0).(func(context.Context) common.ConfigResponse); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(common.ConfigResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context) errors.EdgeX); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// Ping provides a mock function with given fields: ctx
func (_m *CommonClient) Ping(ctx context.Context) (common.PingResponse, errors.EdgeX) {
	ret := _m.Called(ctx)

	var r0 common.PingResponse
	if rf, ok := ret.Get(0).(func(context.Context) common.PingResponse); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(common.PingResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context) errors.EdgeX); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// Version provides a mock function with given fields: ctx
func (_m *CommonClient) Version(ctx context.Context) (common.VersionResponse, errors.EdgeX) {
	ret := _m.Called(ctx)

	var r0 common.VersionResponse
	if rf, ok := ret.Get(0).(func(context.Context) common.VersionResponse); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(common.VersionResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context) errors.EdgeX); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewCommonClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewCommonClient creates a new instance of CommonClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCommonClient(t mockConstructorTestingTNewCommonClient) *CommonClient {
	mock := &CommonClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
