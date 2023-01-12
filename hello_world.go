package http

import "net/http"

// HelloWorld returns a simple HandlerFunc.
func HelloWorld(from string) http.HandlerFunc {
	data := []byte("hello world from " + from)
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(data)
	}
}
