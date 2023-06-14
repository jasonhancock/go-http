package http

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHeaderRoundTripper(t *testing.T) {
	tr := &mockRT{
		RoundTripFunc: func(req *http.Request) (*http.Response, error) {
			require.Contains(t, req.Header, "User-Agent")
			require.Equal(t, "test-1.1", req.Header.Get("User-Agent"))

			return &http.Response{}, nil
		},
	}

	client := &http.Client{
		Transport: UserAgentTransport(tr, "test-1.1"),
	}

	_, err := client.Get("http://foo.example.com")
	require.NoError(t, err)
}

type mockRT struct {
	RoundTripFunc func(req *http.Request) (*http.Response, error)
}

func (m *mockRT) RoundTrip(req *http.Request) (*http.Response, error) {
	if m.RoundTripFunc != nil {
		return m.RoundTripFunc(req)
	}
	panic("not implemented")
}
