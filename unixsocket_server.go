package http

import (
	"context"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

// NewUnixSocketHTTPServer starts up an HTTP server listening over a unix socket.
// The server will run until the context is cancelled.
func NewUnixSocketHTTPServer(ctx context.Context, l Logger, wg *sync.WaitGroup, hler http.Handler, socketFile string) error {
	_ = os.Remove(socketFile)
	unixListener, err := net.Listen("unix", socketFile)
	if err != nil {
		return err
	}

	server := http.Server{
		Handler:           hler,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
	}

	wg.Add(1)
	go func() {
		l.Info("starting http server on unix domain socket", "socket file", socketFile)
		if err := server.Serve(unixListener); err != nil && err != http.ErrServerClosed {
			l.Err("starting http unix socket server error", "error", err.Error())
		}
	}()

	go func() {
		defer wg.Done()
		<-ctx.Done()
		l.Info("stopping http server on unix domain socket", "socket file", socketFile)

		// shut down gracefully, but wait no longer than 10 seconds before halting
		sdCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(sdCtx); err != nil {
			l.Err(
				"stopping http server on unix domain socket error",
				"error", err.Error,
				"socket file", socketFile,
			)
		}
		l.Info(
			"stopped http server on unix domain socket error",
			"socket file", socketFile,
		)
	}()

	return nil
}
