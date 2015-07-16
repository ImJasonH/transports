package logging

import (
	"log"
	"net/http"
)

func NewTransport(t http.RoundTripper) http.RoundTripper {
	return loggingTransport{t}
}

type loggingTransport struct {
	t http.RoundTripper
}

func (lt loggingTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	log.Println(r.Method, r.URL)
	return lt.t.RoundTrip(r)
}
