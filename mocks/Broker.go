// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	broker "github.com/anycable/anycable-go/broker"
	common "github.com/anycable/anycable-go/common"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Broker is an autogenerated mock type for the Broker type
type Broker struct {
	mock.Mock
}

// Announce provides a mock function with given fields:
func (_m *Broker) Announce() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// CommitSession provides a mock function with given fields: sid, session
func (_m *Broker) CommitSession(sid string, session broker.Cacheable) error {
	ret := _m.Called(sid, session)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, broker.Cacheable) error); ok {
		r0 = rf(sid, session)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FinishSession provides a mock function with given fields: sid
func (_m *Broker) FinishSession(sid string) error {
	ret := _m.Called(sid)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(sid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HandleBroadcast provides a mock function with given fields: msg
func (_m *Broker) HandleBroadcast(msg *common.StreamMessage) {
	_m.Called(msg)
}

// HandleCommand provides a mock function with given fields: msg
func (_m *Broker) HandleCommand(msg *common.RemoteCommandMessage) {
	_m.Called(msg)
}

// HistoryFrom provides a mock function with given fields: stream, epoch, offset
func (_m *Broker) HistoryFrom(stream string, epoch string, offset uint64) ([]common.StreamMessage, error) {
	ret := _m.Called(stream, epoch, offset)

	var r0 []common.StreamMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, uint64) ([]common.StreamMessage, error)); ok {
		return rf(stream, epoch, offset)
	}
	if rf, ok := ret.Get(0).(func(string, string, uint64) []common.StreamMessage); ok {
		r0 = rf(stream, epoch, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.StreamMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, uint64) error); ok {
		r1 = rf(stream, epoch, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HistorySince provides a mock function with given fields: stream, ts
func (_m *Broker) HistorySince(stream string, ts int64) ([]common.StreamMessage, error) {
	ret := _m.Called(stream, ts)

	var r0 []common.StreamMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int64) ([]common.StreamMessage, error)); ok {
		return rf(stream, ts)
	}
	if rf, ok := ret.Get(0).(func(string, int64) []common.StreamMessage); ok {
		r0 = rf(stream, ts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.StreamMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int64) error); ok {
		r1 = rf(stream, ts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RestoreSession provides a mock function with given fields: from
func (_m *Broker) RestoreSession(from string) ([]byte, error) {
	ret := _m.Called(from)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]byte, error)); ok {
		return rf(from)
	}
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(from)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(from)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Shutdown provides a mock function with given fields: ctx
func (_m *Broker) Shutdown(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields:
func (_m *Broker) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Subscribe provides a mock function with given fields: stream
func (_m *Broker) Subscribe(stream string) string {
	ret := _m.Called(stream)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(stream)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Unsubscribe provides a mock function with given fields: stream
func (_m *Broker) Unsubscribe(stream string) string {
	ret := _m.Called(stream)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(stream)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewBroker interface {
	mock.TestingT
	Cleanup(func())
}

// NewBroker creates a new instance of Broker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBroker(t mockConstructorTestingTNewBroker) *Broker {
	mock := &Broker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
