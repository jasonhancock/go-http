package http

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetClientIP(t *testing.T) {
	t.Run("without X-Forwarded-For", func(t *testing.T) {
		r := &http.Request{
			RemoteAddr: "192.168.0.1:1234",
		}
		assert.Equal(t, "192.168.0.1", ExtractClientIP(r))
	})

	t.Run("without X-Forwarded-For, unparseable", func(t *testing.T) {
		r := &http.Request{
			RemoteAddr: "foo",
		}
		assert.Equal(t, "", ExtractClientIP(r))
	})

	t.Run("X-Forwarded-For, single entry", func(t *testing.T) {
		r := &http.Request{
			RemoteAddr: "192.168.1.2:1234",
			Header: map[string][]string{
				"X-Forwarded-For": {"192.168.1.1"},
			},
		}
		assert.Equal(t, "192.168.1.1", ExtractClientIP(r))
	})

	t.Run("X-Forwarded-For, multiple entries", func(t *testing.T) {
		r := &http.Request{
			RemoteAddr: "192.168.1.2:1234",
			Header: map[string][]string{
				"X-Forwarded-For": {"192.168.1.4", "192.168.1.1"},
			},
		}
		assert.Equal(t, "192.168.1.1", ExtractClientIP(r))
	})
}
