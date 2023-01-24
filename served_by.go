package http

import (
	"net/http"
	"os"
)

func ServedBy(next http.Handler) http.Handler {
	host, _ := os.Hostname()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Served-By", host)
		next.ServeHTTP(w, r)
	})
}
