// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	common "github.com/agile-edgex/go-mod-core-contracts/v3/dtos/common"

	errors "github.com/agile-edgex/go-mod-core-contracts/v3/errors"

	mock "github.com/stretchr/testify/mock"

	responses "github.com/agile-edgex/go-mod-core-contracts/v3/dtos/responses"
)

// ReadingClient is an autogenerated mock type for the ReadingClient type
type ReadingClient struct {
	mock.Mock
}

// AllReadings provides a mock function with given fields: ctx, offset, limit
func (_m *ReadingClient) AllReadings(ctx context.Context, offset int, limit int) (responses.MultiReadingsResponse, errors.EdgeX) {
	ret := _m.Called(ctx, offset, limit)

	var r0 responses.MultiReadingsResponse
	if rf, ok := ret.Get(0).(func(context.Context, int, int) responses.MultiReadingsResponse); ok {
		r0 = rf(ctx, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiReadingsResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ReadingCount provides a mock function with given fields: ctx
func (_m *ReadingClient) ReadingCount(ctx context.Context) (common.CountResponse, errors.EdgeX) {
	ret := _m.Called(ctx)

	var r0 common.CountResponse
	if rf, ok := ret.Get(0).(func(context.Context) common.CountResponse); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(common.CountResponse)
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

// ReadingCountByDeviceName provides a mock function with given fields: ctx, name
func (_m *ReadingClient) ReadingCountByDeviceName(ctx context.Context, name string) (common.CountResponse, errors.EdgeX) {
	ret := _m.Called(ctx, name)

	var r0 common.CountResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) common.CountResponse); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(common.CountResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string) errors.EdgeX); ok {
		r1 = rf(ctx, name)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ReadingsByDeviceName provides a mock function with given fields: ctx, name, offset, limit
func (_m *ReadingClient) ReadingsByDeviceName(ctx context.Context, name string, offset int, limit int) (responses.MultiReadingsResponse, errors.EdgeX) {
	ret := _m.Called(ctx, name, offset, limit)

	var r0 responses.MultiReadingsResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) responses.MultiReadingsResponse); ok {
		r0 = rf(ctx, name, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiReadingsResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, name, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ReadingsByDeviceNameAndResourceName provides a mock function with given fields: ctx, deviceName, resourceName, offset, limit
func (_m *ReadingClient) ReadingsByDeviceNameAndResourceName(ctx context.Context, deviceName string, resourceName string, offset int, limit int) (responses.MultiReadingsResponse, errors.EdgeX) {
	ret := _m.Called(ctx, deviceName, resourceName, offset, limit)

	var r0 responses.MultiReadingsResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int, int) responses.MultiReadingsResponse); ok {
		r0 = rf(ctx, deviceName, resourceName, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiReadingsResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string, string, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, deviceName, resourceName, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ReadingsByDeviceNameAndResourceNameAndTimeRange provides a mock function with given fields: ctx, deviceName, resourceName, start, end, offset, limit
func (_m *ReadingClient) ReadingsByDeviceNameAndResourceNameAndTimeRange(ctx context.Context, deviceName string, resourceName string, start int, end int, offset int, limit int) (responses.MultiReadingsResponse, errors.EdgeX) {
	ret := _m.Called(ctx, deviceName, resourceName, start, end, offset, limit)

	var r0 responses.MultiReadingsResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int, int, int, int) responses.MultiReadingsResponse); ok {
		r0 = rf(ctx, deviceName, resourceName, start, end, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiReadingsResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string, string, int, int, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, deviceName, resourceName, start, end, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ReadingsByDeviceNameAndResourceNamesAndTimeRange provides a mock function with given fields: ctx, deviceName, resourceNames, start, end, offset, limit
func (_m *ReadingClient) ReadingsByDeviceNameAndResourceNamesAndTimeRange(ctx context.Context, deviceName string, resourceNames []string, start int, end int, offset int, limit int) (responses.MultiReadingsResponse, errors.EdgeX) {
	ret := _m.Called(ctx, deviceName, resourceNames, start, end, offset, limit)

	var r0 responses.MultiReadingsResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, []string, int, int, int, int) responses.MultiReadingsResponse); ok {
		r0 = rf(ctx, deviceName, resourceNames, start, end, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiReadingsResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string, []string, int, int, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, deviceName, resourceNames, start, end, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ReadingsByResourceName provides a mock function with given fields: ctx, name, offset, limit
func (_m *ReadingClient) ReadingsByResourceName(ctx context.Context, name string, offset int, limit int) (responses.MultiReadingsResponse, errors.EdgeX) {
	ret := _m.Called(ctx, name, offset, limit)

	var r0 responses.MultiReadingsResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) responses.MultiReadingsResponse); ok {
		r0 = rf(ctx, name, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiReadingsResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, name, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ReadingsByResourceNameAndTimeRange provides a mock function with given fields: ctx, name, start, end, offset, limit
func (_m *ReadingClient) ReadingsByResourceNameAndTimeRange(ctx context.Context, name string, start int, end int, offset int, limit int) (responses.MultiReadingsResponse, errors.EdgeX) {
	ret := _m.Called(ctx, name, start, end, offset, limit)

	var r0 responses.MultiReadingsResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int, int, int) responses.MultiReadingsResponse); ok {
		r0 = rf(ctx, name, start, end, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiReadingsResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, string, int, int, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, name, start, end, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

// ReadingsByTimeRange provides a mock function with given fields: ctx, start, end, offset, limit
func (_m *ReadingClient) ReadingsByTimeRange(ctx context.Context, start int, end int, offset int, limit int) (responses.MultiReadingsResponse, errors.EdgeX) {
	ret := _m.Called(ctx, start, end, offset, limit)

	var r0 responses.MultiReadingsResponse
	if rf, ok := ret.Get(0).(func(context.Context, int, int, int, int) responses.MultiReadingsResponse); ok {
		r0 = rf(ctx, start, end, offset, limit)
	} else {
		r0 = ret.Get(0).(responses.MultiReadingsResponse)
	}

	var r1 errors.EdgeX
	if rf, ok := ret.Get(1).(func(context.Context, int, int, int, int) errors.EdgeX); ok {
		r1 = rf(ctx, start, end, offset, limit)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.EdgeX)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewReadingClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewReadingClient creates a new instance of ReadingClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReadingClient(t mockConstructorTestingTNewReadingClient) *ReadingClient {
	mock := &ReadingClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
