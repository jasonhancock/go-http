package http

type contextKey int

const (
	keyClientIP contextKey = iota
	keyRequestID
)
