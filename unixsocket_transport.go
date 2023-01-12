package http

import (
	"context"
	"net"
	"net/http"

	"github.com/hashicorp/go-cleanhttp"
)

// NewUnixSocketRoundTripper sets up an http RoundTripper that will communicate
// over a unix domain socket. The baseTransport is optional.
func NewUnixSocketRoundTripper(socketFile string, baseTransport *http.Transport) *UnixRoundTripper {
	if baseTransport == nil {
		baseTransport = cleanhttp.DefaultTransport()
	}

	baseTransport.DialContext = func(_ context.Context, _, _ string) (net.Conn, error) {
		return net.Dial("unix", socketFile)
	}

	return &UnixRoundTripper{
		rt: baseTransport,
	}
}

// UnixRoundTripper is an http.RoundTripper that makes sure the scheme and host
// is set correctly to talk over a unix socket.
type UnixRoundTripper struct {
	rt http.RoundTripper
}

func (rt *UnixRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	request.URL.Scheme = "http"
	request.URL.Host = "unix"

	return rt.rt.RoundTrip(request)
}
