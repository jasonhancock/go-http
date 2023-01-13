package http

import "io"

// Drain finishes reading a response body and closes it.
func Drain(r io.ReadCloser) error {
	if _, err := io.Copy(io.Discard, r); err != nil {
		return err
	}
	return r.Close()
}
