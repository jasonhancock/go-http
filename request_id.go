package http

import (
	"context"
	"net/http"

	"github.com/google/uuid"
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
