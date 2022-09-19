package geocodingsearchv7_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"go.einride.tech/here/geocodingsearchv7"
	"gotest.tools/v3/assert"
)

type ReverseGeocodingMock struct {
	responseStatus int
	responseBody   geocodingsearchv7.ReverseGeocodingResponse
}

func (c *ReverseGeocodingMock) Do(req *http.Request) (*http.Response, error) {
	headers := http.Header{}
	headers.Add("Content-Type", "application/json")
	b, err := json.Marshal(c.responseBody)
	if err != nil {
		return nil, err
	}
	r := bytes.NewReader(b)
	return &http.Response{
		StatusCode:    c.responseStatus,
		Header:        headers,
		Body:          io.NopCloser(r),
		ContentLength: int64(len(b)),
	}, nil
}

func TestReverseGeocodingService_ReverseGeocoding(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run(
		"when geoposition is missing from request, then return invalid argument",
		func(t *testing.T) {
			t.Parallel()
			httpClient := ReverseGeocodingMock{responseBody: geocodingsearchv7.ReverseGeocodingResponse{}, responseStatus: 200}
			routingClient := geocodingsearchv7.NewClient(&httpClient)

			_, err := routingClient.ReverseGeocoding.ReverseGeocoding(ctx, &geocodingsearchv7.ReverseGeocodingRequest{})
			assert.ErrorContains(t, err, "InvalidArgument")
		},
	)
}
