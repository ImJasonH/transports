// Package fn provides a function to create an http.RoundTripper from a func.
package fn

import "net/http"

// RoundTripperFunc returns an http.RoundTripper that executes the given func.
func RoundTripperFunc(f func(*http.Request) (*http.Response, error)) http.RoundTripper {
	return rtFunc(f)
}

type rtFunc func(*http.Request) (*http.Response, error)

func (rtFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return rtFunc(r)
}
