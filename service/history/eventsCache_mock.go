// The MIT License (MIT)
//
// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: eventsCache.go

// Package history is a generated GoMock package.
package history

import (
	gomock "github.com/golang/mock/gomock"
	shared "github.com/uber/cadence/.gen/go/shared"
	reflect "reflect"
)

// MockeventsCache is a mock of eventsCache interface
type MockeventsCache struct {
	ctrl     *gomock.Controller
	recorder *MockeventsCacheMockRecorder
}

// MockeventsCacheMockRecorder is the mock recorder for MockeventsCache
type MockeventsCacheMockRecorder struct {
	mock *MockeventsCache
}

// NewMockeventsCache creates a new mock instance
func NewMockeventsCache(ctrl *gomock.Controller) *MockeventsCache {
	mock := &MockeventsCache{ctrl: ctrl}
	mock.recorder = &MockeventsCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockeventsCache) EXPECT() *MockeventsCacheMockRecorder {
	return m.recorder
}

// getEvent mocks base method
func (m *MockeventsCache) getEvent(domainID, workflowID, runID string, firstEventID, eventID int64, branchToken []byte) (*shared.HistoryEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getEvent", domainID, workflowID, runID, firstEventID, eventID, branchToken)
	ret0, _ := ret[0].(*shared.HistoryEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getEvent indicates an expected call of getEvent
func (mr *MockeventsCacheMockRecorder) getEvent(domainID, workflowID, runID, firstEventID, eventID, branchToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getEvent", reflect.TypeOf((*MockeventsCache)(nil).getEvent), domainID, workflowID, runID, firstEventID, eventID, branchToken)
}

// putEvent mocks base method
func (m *MockeventsCache) putEvent(domainID, workflowID, runID string, eventID int64, event *shared.HistoryEvent) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "putEvent", domainID, workflowID, runID, eventID, event)
}

// putEvent indicates an expected call of putEvent
func (mr *MockeventsCacheMockRecorder) putEvent(domainID, workflowID, runID, eventID, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "putEvent", reflect.TypeOf((*MockeventsCache)(nil).putEvent), domainID, workflowID, runID, eventID, event)
}

// deleteEvent mocks base method
func (m *MockeventsCache) deleteEvent(domainID, workflowID, runID string, eventID int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "deleteEvent", domainID, workflowID, runID, eventID)
}

// deleteEvent indicates an expected call of deleteEvent
func (mr *MockeventsCacheMockRecorder) deleteEvent(domainID, workflowID, runID, eventID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "deleteEvent", reflect.TypeOf((*MockeventsCache)(nil).deleteEvent), domainID, workflowID, runID, eventID)
}
