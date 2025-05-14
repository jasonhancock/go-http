package http

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStdlibLogger(t *testing.T) {
	t.Run("non TLS handshake error", func(t *testing.T) {
		logger := LoggerMock{
			ErrFunc:   func(msg any, keyvals ...any) {},
			DebugFunc: func(msg any, keyvals ...any) {},
		}
		l := stdlibLogger(&logger)

		l.Println("foo")
		l.Println("TLS handshake error")

		require.Len(t, logger.ErrCalls(), 2)
		require.Len(t, logger.DebugCalls(), 0)
	})
}

func TestTLSHandshakeFilteringLogger(t *testing.T) {
	t.Run("non TLS handshake error", func(t *testing.T) {
		logger := LoggerMock{
			ErrFunc:   func(msg any, keyvals ...any) {},
			DebugFunc: func(msg any, keyvals ...any) {},
		}
		l := TLSHandshakeFilteringLogger(&logger)

		l.Println("foo")

		require.Len(t, logger.ErrCalls(), 1)
		require.Len(t, logger.DebugCalls(), 0)
	})

	t.Run("TLS handshake error", func(t *testing.T) {
		logger := LoggerMock{
			ErrFunc:   func(msg any, keyvals ...any) {},
			DebugFunc: func(msg any, keyvals ...any) {},
		}
		l := TLSHandshakeFilteringLogger(&logger)

		l.Println("TLS handshake error")

		require.Len(t, logger.ErrCalls(), 0)
		require.Len(t, logger.DebugCalls(), 1)
	})
}
