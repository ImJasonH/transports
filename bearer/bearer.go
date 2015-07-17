// Package bearer provides an http.RoundTripper that sets a bearer
// authorization token on all requests.
package bearer

import "net/http"

// NewTransport returns a RoundTripper that sets the provided token as a bearer
// authorization token on all requests.
func NewTransport(t http.RoundTripper, tok string) http.RoundTripper {
	return bearerTransport{t, tok}
}

type bearerTransport struct {
	t   http.RoundTripper
	tok string
}

func (bt bearerTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("Authorization", "Bearer "+bt.tok)
	return bt.t.RoundTrip(r)
}
