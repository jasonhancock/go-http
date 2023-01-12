package http

import (
	"context"
	"net"
	"net/http"
)

// ExtractClientIP returns the client IP from the given http.Request. Respects the
// X-Forwarded-For header.
func ExtractClientIP(r *http.Request) string {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)

	forwardedValues, ok := r.Header["X-Forwarded-For"]
	if !ok {
		return ip
	}

	// the client IP is the last in the list of X-Forwarded-For
	l := len(forwardedValues)
	if l > 0 {
		ip = forwardedValues[l-1]
	}

	return ip
}

type clientIP int

const clientIPKey clientIP = iota

// ClientIP is middleware that adds the client IP address to the context in the request.
func ClientIP(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r.WithContext(WithClientIP(r.Context(), ExtractClientIP(r))))
	}
	return http.HandlerFunc(fn)
}

// WithClientIP sets the IP in the context.
func WithClientIP(ctx context.Context, ip string) context.Context {
	return context.WithValue(ctx, clientIPKey, ip)
}

// GetClientIP returns the client IP from the context.
func GetClientIP(ctx context.Context) string {
	s, _ := ctx.Value(clientIPKey).(string)
	return s
}
