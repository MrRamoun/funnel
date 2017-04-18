package mocks

import funnel "funnel/proto/funnel"

import mock "github.com/stretchr/testify/mock"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// StartWorker provides a mock function with given fields: tplName, serverAddress, workerID
func (_m *Client) StartWorker(tplName string, serverAddress string, workerID string) error {
	ret := _m.Called(tplName, serverAddress, workerID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(tplName, serverAddress, workerID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Templates provides a mock function with given fields:
func (_m *Client) Templates() []funnel.Worker {
	ret := _m.Called()

	var r0 []funnel.Worker
	if rf, ok := ret.Get(0).(func() []funnel.Worker); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]funnel.Worker)
		}
	}

	return r0
}
