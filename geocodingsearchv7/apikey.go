package geocodingsearchv7

import "net/http"

type apiKeyRoundTripper struct {
	apiKey string
	next   http.RoundTripper
}

// NewAPIKeyHTTPClient returns an HTTP Client which uses the given API Key.
// If next is nil http.DefaultTransport is used.
func NewAPIKeyHTTPClient(key string, next http.RoundTripper) *http.Client {
	return &http.Client{
		Transport: &apiKeyRoundTripper{
			apiKey: key,
			next:   next,
		},
	}
}

func (r *apiKeyRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	vals := req.URL.Query()
	vals.Set("apiKey", r.apiKey)
	req.URL.RawQuery = vals.Encode()
	if r.next != nil {
		return r.next.RoundTrip(req)
	}
	return http.DefaultTransport.RoundTrip(req)
}
