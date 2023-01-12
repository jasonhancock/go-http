package http

// IsOK returns true if the response code is a 2xx.
func IsOK(code int) bool {
	return code >= 200 && code <= 299
}
