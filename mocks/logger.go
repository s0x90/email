// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"sync"
)

// LoggerMock is a mock implementation of notify.Logger.
//
// 	func TestSomethingThatUsesLogger(t *testing.T) {
//
// 		// make and configure a mocked notify.Logger
// 		mockedLogger := &LoggerMock{
// 			LogfFunc: func(format string, args ...interface{})  {
// 				panic("mock out the Logf method")
// 			},
// 		}
//
// 		// use mockedLogger in code that requires notify.Logger
// 		// and then make assertions.
//
// 	}
type LoggerMock struct {
	// LogfFunc mocks the Logf method.
	LogfFunc func(format string, args ...interface{})

	// calls tracks calls to the methods.
	calls struct {
		// Logf holds details about calls to the Logf method.
		Logf []struct {
			// Format is the format argument value.
			Format string
			// Args is the args argument value.
			Args []interface{}
		}
	}
	lockLogf sync.RWMutex
}

// Logf calls LogfFunc.
func (mock *LoggerMock) Logf(format string, args ...interface{}) {
	if mock.LogfFunc == nil {
		panic("LoggerMock.LogfFunc: method is nil but Logger.Logf was just called")
	}
	callInfo := struct {
		Format string
		Args   []interface{}
	}{
		Format: format,
		Args:   args,
	}
	mock.lockLogf.Lock()
	mock.calls.Logf = append(mock.calls.Logf, callInfo)
	mock.lockLogf.Unlock()
	mock.LogfFunc(format, args...)
}

// LogfCalls gets all the calls that were made to Logf.
// Check the length with:
//     len(mockedLogger.LogfCalls())
func (mock *LoggerMock) LogfCalls() []struct {
	Format string
	Args   []interface{}
} {
	var calls []struct {
		Format string
		Args   []interface{}
	}
	mock.lockLogf.RLock()
	calls = mock.calls.Logf
	mock.lockLogf.RUnlock()
	return calls
}