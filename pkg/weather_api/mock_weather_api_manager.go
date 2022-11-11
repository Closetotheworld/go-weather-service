// Code generated by mockery v2.14.1. DO NOT EDIT.

package weather_api

import mock "github.com/stretchr/testify/mock"

// MockWeatherApiManager is an autogenerated mock type for the WeatherApiManager type
type MockWeatherApiManager struct {
	mock.Mock
}

// AsyncRequest provides a mock function with given fields: lat, lon
func (_m *MockWeatherApiManager) AsyncRequest(lat string, lon string) (*WeatherApiCommon, []*WeatherApiForecast, []*WeatherApiCommon, error) {
	ret := _m.Called(lat, lon)

	var r0 *WeatherApiCommon
	if rf, ok := ret.Get(0).(func(string, string) *WeatherApiCommon); ok {
		r0 = rf(lat, lon)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*WeatherApiCommon)
		}
	}

	var r1 []*WeatherApiForecast
	if rf, ok := ret.Get(1).(func(string, string) []*WeatherApiForecast); ok {
		r1 = rf(lat, lon)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*WeatherApiForecast)
		}
	}

	var r2 []*WeatherApiCommon
	if rf, ok := ret.Get(2).(func(string, string) []*WeatherApiCommon); ok {
		r2 = rf(lat, lon)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).([]*WeatherApiCommon)
		}
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(string, string) error); ok {
		r3 = rf(lat, lon)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

type mockConstructorTestingTNewMockWeatherApiManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockWeatherApiManager creates a new instance of MockWeatherApiManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockWeatherApiManager(t mockConstructorTestingTNewMockWeatherApiManager) *MockWeatherApiManager {
	mock := &MockWeatherApiManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}