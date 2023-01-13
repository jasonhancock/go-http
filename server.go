package http

import (
	"context"
	"crypto/tls"
	"net/http"
	"sync"
	"time"
)

// Logger is an interface for logging messages.
type Logger interface {
	Err(msg interface{}, keyvals ...interface{})
	Info(msg interface{}, keyvals ...interface{})
	Fatal(msg interface{}, keyvals ...interface{})
}

type options struct {
	TLSConfig         *tls.Config
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	ReadHeaderTimeout time.Duration
}

// ServerOption is used to customize the server.
type ServerOption func(*options)

// WithTLSConfig sets the TLS configuration to use on the server.
func WithTLSConfig(c *tls.Config) ServerOption {
	return func(o *options) {
		o.TLSConfig = c
	}
}

// WithTimeouts sets the server's ReadTimeout, WriteTimeout, and ReadHeaderTimeout.
func WithTimeeouts(t time.Duration) ServerOption {
	return func(o *options) {
		o.ReadTimeout = t
		o.WriteTimeout = t
		o.ReadHeaderTimeout = t
	}
}

// NewHTTPServer starts up an HTTP server. The server will run until the context
// is cancelled.
func NewHTTPServer(ctx context.Context, l Logger, wg *sync.WaitGroup, hler http.Handler, addr string, opts ...ServerOption) {
	opt := options{
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
	}
	for _, o := range opts {
		o(&opt)
	}

	server := http.Server{
		Addr:              addr,
		Handler:           hler,
		ReadTimeout:       opt.ReadTimeout,
		WriteTimeout:      opt.WriteTimeout,
		ReadHeaderTimeout: opt.ReadHeaderTimeout,
	}

	wg.Add(1)
	go func() {
		if opt.TLSConfig == nil {
			l.Info("starting http server", "addr", addr)
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				l.Fatal("starting http server error", "error", err.Error(), "addr", addr)
			}
		} else {
			server.Handler = StrictHSTS(server.Handler)
			server.TLSConfig = opt.TLSConfig
			// This next line disables HTTP/2 because it mandates the use of a weak 128 bit cipher.
			server.TLSNextProto = make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0)
			l.Info("starting https server", "addr", addr)
			if err := server.ListenAndServeTLS("", ""); err != nil && err != http.ErrServerClosed {
				l.Fatal("starting https server error", "error", err.Error(), "addr", addr)
			}
		}
	}()

	go func() {
		defer wg.Done()
		<-ctx.Done()
		l.Info("stopping http server", "addr", addr)

		// shut down gracefully, but wait no longer than 10 seconds before halting
		sdCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(sdCtx); err != nil {
			l.Err(
				"stopping http server",
				"error", err.Error(),
				"addr", addr,
			)
		}
		l.Info("stopped http server", "addr", addr)
	}()
}

// HardenedTLSConfig returns a tls.Config with the basics configured. It follows
// this guide: https://blog.bracebin.com/achieving-perfect-ssl-labs-score-with-go
func HardenedTLSConfig() *tls.Config {
	// #nosec G402
	return &tls.Config{
		MinVersion:               tls.VersionTLS12,
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
		},
		CurvePreferences: []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
	}
}

func StrictHSTS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		next.ServeHTTP(w, r)
	})
}
