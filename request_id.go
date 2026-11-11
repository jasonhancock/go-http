package http

import (
	"context"
	"net/http"
	"uuid"
)

// RequestHeader is the header that marks the http request id set by the server.
const RequestHeader = "X-Http-Request-Id"

// RequestID sets the request id in the context to a UUID that can be traced
// throughout the servers req/resp lifecycle.
func RequestID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set(RequestHeader, id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), keyRequestID, id)))
	}
	return http.HandlerFunc(fn)
}

// GetRequestID returns any request id that had been set in the context.
func GetRequestID(ctx context.Context) string {
	s, _ := ctx.Value(keyRequestID).(string)
	return s
}

// ResolveRequestID returns the request id set in ctx by the RequestID
// middleware as a "request_id" key/value pair, or nil if no request id is set.
//
// It satisfies github.com/jasonhancock/go-logger's ContextResolutionFunc, so
// the request id can be added automatically to log lines written with the
// logger's *Context methods:
//
//	l := logger.New(logger.WithContextResolutionFunc(http.ResolveRequestID))
//	l.InfoContext(r.Context(), "handling request")
func ResolveRequestID(ctx context.Context) []any {
	id, ok := ctx.Value(keyRequestID).(string)
	if !ok || id == "" {
		return nil
	}

	return []any{"request_id", id}
}
