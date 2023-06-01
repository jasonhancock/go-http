package http

import "log"

// stdlibLogger initializes a *log.Logger that can be used as a net/http.Server's ErrorLog.
func stdlibLogger(l Logger) *log.Logger {
	return log.New(&logAdapter{logger: l}, "", 0)
}

type logAdapter struct {
	logger Logger
}

func (l *logAdapter) Write(data []byte) (int, error) {
	l.logger.Err(string(data))
	return len(data), nil
}
