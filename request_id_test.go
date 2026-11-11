package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

// ResolveRequestID must keep the signature of go-logger's
// ContextResolutionFunc so it can be passed to WithContextResolutionFunc.
var _ func(context.Context) []any = ResolveRequestID

func TestResolveRequestID(t *testing.T) {
	tests := []struct {
		name string
		ctx  context.Context
		want []any
	}{
		{
			name: "no request id",
			ctx:  context.Background(),
			want: nil,
		},
		{
			name: "empty request id",
			ctx:  context.WithValue(context.Background(), keyRequestID, ""),
			want: nil,
		},
		{
			name: "wrong type",
			ctx:  context.WithValue(context.Background(), keyRequestID, 1234),
			want: nil,
		},
		{
			name: "request id set",
			ctx:  context.WithValue(context.Background(), keyRequestID, "abc123"),
			want: []any{"request_id", "abc123"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, ResolveRequestID(tt.ctx))
		})
	}
}

func TestResolveRequestIDWithMiddleware(t *testing.T) {
	var got []any
	h := RequestID(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = ResolveRequestID(r.Context())
	}))

	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))

	id := rec.Header().Get(RequestHeader)
	require.NotEmpty(t, id)
	require.Equal(t, []any{"request_id", id}, got)
}
