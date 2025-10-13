package http

import (
	"log"
	"strings"
)

// StdlibLogger initializes a *log.Logger that can be used as a net/http.Server's ErrorLog.
func StdlibLogger(l Logger) *log.Logger {
	return log.New(&logAdapter{logger: l}, "", 0)
}

type logAdapter struct {
	logger Logger
}

func (l *logAdapter) Write(data []byte) (int, error) {
	l.logger.Err(string(data))
	return len(data), nil
}

// TLSHandshakeFilteringLogger routes TLS handshake errors to the Debug log
// level. This can help mitigate log spam.
func TLSHandshakeFilteringLogger(l Logger) *log.Logger {
	return log.New(&tlsHandshakeFilteringLogAdapter{logger: l}, "", 0)
}

type tlsHandshakeFilteringLogAdapter struct {
	logger Logger
}

func (l *tlsHandshakeFilteringLogAdapter) Write(data []byte) (int, error) {
	str := string(data)

	var fn func(msg any, keyvals ...any) = l.logger.Err

	if strings.Contains(str, "TLS handshake error") {
		fn = l.logger.Debug
	}
	fn(str)
	return len(data), nil
}
