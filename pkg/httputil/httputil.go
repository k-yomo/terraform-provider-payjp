package httputil

import "net/http"

type AddHeaderTransport struct {
	transport http.RoundTripper
	headers   map[string]string
}

func NewAddHeaderTransport(transport http.RoundTripper, headers map[string]string) *AddHeaderTransport {
	if transport == nil {
		transport = http.DefaultTransport
	}
	return &AddHeaderTransport{transport: transport, headers: headers}
}
func (adt *AddHeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range adt.headers {
		req.Header.Add(k, v)
	}
	return adt.transport.RoundTrip(req)
}
