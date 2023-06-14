package http

import "net/http"

// HeaderTransport is an http.RoundTripper that adds a header to each request.
type HeaderTransport struct {
	transport http.RoundTripper
	key       string
	value     string
}

// NewHeaderTransport sets up a transport that wraps the provided transport.
func NewHeaderTransport(t http.RoundTripper, key, value string) *HeaderTransport {
	return &HeaderTransport{
		transport: t,
		key:       key,
		value:     value,
	}
}

// RoundTrip adds the header to the request and passes it along to the
// wrapped transport.
func (t *HeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set(t.key, t.value)
	return t.transport.RoundTrip(req)
}

// UserAgentTransport sets the User-Agent on all requests.
func UserAgentTransport(t http.RoundTripper, userAgent string) http.RoundTripper {
	return NewHeaderTransport(t, "User-Agent", userAgent)
}
