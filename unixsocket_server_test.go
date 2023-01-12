package http

import (
	"context"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnixSocket(t *testing.T) {
	dir, err := os.MkdirTemp("", "")
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, os.RemoveAll(dir))
	})

	r := http.NewServeMux()
	r.HandleFunc("/", handle)

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(func() {
		cancel()
	})
	var wg sync.WaitGroup

	socketFile := filepath.Join(dir, "socket")

	require.NoError(t, NewUnixSocketHTTPServer(ctx, &nullLogger{}, &wg, r, socketFile))

	client := &http.Client{
		Transport: NewUnixSocketRoundTripper(socketFile, nil),
	}

	t.Run("get request", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/foo", nil)
		require.NoError(t, err)

		resp, err := client.Do(req)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		defer resp.Body.Close()
		b, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.Equal(t, "hello world", string(b))
	})

	t.Run("post request", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/foo", strings.NewReader("some body"))
		require.NoError(t, err)

		resp, err := client.Do(req)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		defer resp.Body.Close()
		b, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.Equal(t, "hello world", string(b))
	})

	cancel()
	wg.Wait()

	// Make sure the server cleaned up the socket file
	_, err = os.Stat(socketFile)
	require.Error(t, err)
	require.True(t, os.IsNotExist(err))

}
