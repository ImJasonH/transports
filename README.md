This repo provides HTTP RoundTrippers I find myself reimplementing all the time.

bearer
===

[![GoDoc](https://godoc.org/github.com/ImJasonH/transports/bearer?status.svg)](https://godoc.org/github.com/ImJasonH/transports/bearer)

This package provides an `http.RoundTripper` that sets a bearer authorization
token on all requests.

logging
===

[![GoDoc](https://godoc.org/github.com/ImJasonH/transports/logging?status.svg)](https://godoc.org/github.com/ImJasonH/transports/logging)

This package provides an `http.RoundTripper` that logs all requests it makes.

fn
===

[![GoDoc](https://godoc.org/github.com/ImJasonH/transports/fn?status.svg)](https://godoc.org/github.com/ImJasonH/transports/fn)

This package provides a func to create an `http.RoundTripper` from a func.

```
client := &http.Client{
  Transport: fn.RoundTripperFunc(func(r *http.Request) (*http.Response, error) {
    // this func is called when the client RoundTrips.
  })
}
```
