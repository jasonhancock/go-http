// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package http

import (
	"sync"
)

// Ensure, that LoggerMock does implement Logger.
// If this is not the case, regenerate this file with moq.
var _ Logger = &LoggerMock{}

// LoggerMock is a mock implementation of Logger.
//
//	func TestSomethingThatUsesLogger(t *testing.T) {
//
//		// make and configure a mocked Logger
//		mockedLogger := &LoggerMock{
//			DebugFunc: func(msg any, keyvals ...any)  {
//				panic("mock out the Debug method")
//			},
//			ErrFunc: func(msg any, keyvals ...any)  {
//				panic("mock out the Err method")
//			},
//			FatalFunc: func(msg any, keyvals ...any)  {
//				panic("mock out the Fatal method")
//			},
//			InfoFunc: func(msg any, keyvals ...any)  {
//				panic("mock out the Info method")
//			},
//		}
//
//		// use mockedLogger in code that requires Logger
//		// and then make assertions.
//
//	}
type LoggerMock struct {
	// DebugFunc mocks the Debug method.
	DebugFunc func(msg any, keyvals ...any)

	// ErrFunc mocks the Err method.
	ErrFunc func(msg any, keyvals ...any)

	// FatalFunc mocks the Fatal method.
	FatalFunc func(msg any, keyvals ...any)

	// InfoFunc mocks the Info method.
	InfoFunc func(msg any, keyvals ...any)

	// calls tracks calls to the methods.
	calls struct {
		// Debug holds details about calls to the Debug method.
		Debug []struct {
			// Msg is the msg argument value.
			Msg any
			// Keyvals is the keyvals argument value.
			Keyvals []any
		}
		// Err holds details about calls to the Err method.
		Err []struct {
			// Msg is the msg argument value.
			Msg any
			// Keyvals is the keyvals argument value.
			Keyvals []any
		}
		// Fatal holds details about calls to the Fatal method.
		Fatal []struct {
			// Msg is the msg argument value.
			Msg any
			// Keyvals is the keyvals argument value.
			Keyvals []any
		}
		// Info holds details about calls to the Info method.
		Info []struct {
			// Msg is the msg argument value.
			Msg any
			// Keyvals is the keyvals argument value.
			Keyvals []any
		}
	}
	lockDebug sync.RWMutex
	lockErr   sync.RWMutex
	lockFatal sync.RWMutex
	lockInfo  sync.RWMutex
}

// Debug calls DebugFunc.
func (mock *LoggerMock) Debug(msg any, keyvals ...any) {
	if mock.DebugFunc == nil {
		panic("LoggerMock.DebugFunc: method is nil but Logger.Debug was just called")
	}
	callInfo := struct {
		Msg     any
		Keyvals []any
	}{
		Msg:     msg,
		Keyvals: keyvals,
	}
	mock.lockDebug.Lock()
	mock.calls.Debug = append(mock.calls.Debug, callInfo)
	mock.lockDebug.Unlock()
	mock.DebugFunc(msg, keyvals...)
}

// DebugCalls gets all the calls that were made to Debug.
// Check the length with:
//
//	len(mockedLogger.DebugCalls())
func (mock *LoggerMock) DebugCalls() []struct {
	Msg     any
	Keyvals []any
} {
	var calls []struct {
		Msg     any
		Keyvals []any
	}
	mock.lockDebug.RLock()
	calls = mock.calls.Debug
	mock.lockDebug.RUnlock()
	return calls
}

// Err calls ErrFunc.
func (mock *LoggerMock) Err(msg any, keyvals ...any) {
	if mock.ErrFunc == nil {
		panic("LoggerMock.ErrFunc: method is nil but Logger.Err was just called")
	}
	callInfo := struct {
		Msg     any
		Keyvals []any
	}{
		Msg:     msg,
		Keyvals: keyvals,
	}
	mock.lockErr.Lock()
	mock.calls.Err = append(mock.calls.Err, callInfo)
	mock.lockErr.Unlock()
	mock.ErrFunc(msg, keyvals...)
}

// ErrCalls gets all the calls that were made to Err.
// Check the length with:
//
//	len(mockedLogger.ErrCalls())
func (mock *LoggerMock) ErrCalls() []struct {
	Msg     any
	Keyvals []any
} {
	var calls []struct {
		Msg     any
		Keyvals []any
	}
	mock.lockErr.RLock()
	calls = mock.calls.Err
	mock.lockErr.RUnlock()
	return calls
}

// Fatal calls FatalFunc.
func (mock *LoggerMock) Fatal(msg any, keyvals ...any) {
	if mock.FatalFunc == nil {
		panic("LoggerMock.FatalFunc: method is nil but Logger.Fatal was just called")
	}
	callInfo := struct {
		Msg     any
		Keyvals []any
	}{
		Msg:     msg,
		Keyvals: keyvals,
	}
	mock.lockFatal.Lock()
	mock.calls.Fatal = append(mock.calls.Fatal, callInfo)
	mock.lockFatal.Unlock()
	mock.FatalFunc(msg, keyvals...)
}

// FatalCalls gets all the calls that were made to Fatal.
// Check the length with:
//
//	len(mockedLogger.FatalCalls())
func (mock *LoggerMock) FatalCalls() []struct {
	Msg     any
	Keyvals []any
} {
	var calls []struct {
		Msg     any
		Keyvals []any
	}
	mock.lockFatal.RLock()
	calls = mock.calls.Fatal
	mock.lockFatal.RUnlock()
	return calls
}

// Info calls InfoFunc.
func (mock *LoggerMock) Info(msg any, keyvals ...any) {
	if mock.InfoFunc == nil {
		panic("LoggerMock.InfoFunc: method is nil but Logger.Info was just called")
	}
	callInfo := struct {
		Msg     any
		Keyvals []any
	}{
		Msg:     msg,
		Keyvals: keyvals,
	}
	mock.lockInfo.Lock()
	mock.calls.Info = append(mock.calls.Info, callInfo)
	mock.lockInfo.Unlock()
	mock.InfoFunc(msg, keyvals...)
}

// InfoCalls gets all the calls that were made to Info.
// Check the length with:
//
//	len(mockedLogger.InfoCalls())
func (mock *LoggerMock) InfoCalls() []struct {
	Msg     any
	Keyvals []any
} {
	var calls []struct {
		Msg     any
		Keyvals []any
	}
	mock.lockInfo.RLock()
	calls = mock.calls.Info
	mock.lockInfo.RUnlock()
	return calls
}
