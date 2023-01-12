package http

import "net/http"

// Doer is an interface the http.Client conforms to.
type Doer interface {
	Do(*http.Request) (*http.Response, error)
}
