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

type GeocodingMock struct {
	responseStatus int
	responseBody   geocodingsearchv7.GeocodingResponse
}

func (c *GeocodingMock) Do(req *http.Request) (*http.Response, error) {
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

func TestGeocodingService_Geocoding(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run(
		"when neither address or query is given as input, then return invalid argument",
		func(t *testing.T) {
			t.Parallel()
			httpClient := GeocodingMock{responseBody: geocodingsearchv7.GeocodingResponse{}, responseStatus: 200}
			routingClient := geocodingsearchv7.NewClient(&httpClient)

			_, err := routingClient.Geocoding.Geocoding(ctx, &geocodingsearchv7.GeocodingRequest{
				Address: nil, // Address is nil
				Q:       nil, // Query is nil
			})
			assert.ErrorContains(t, err, "InvalidArgument")
		},
	)
}

func TestGeocodingService_FormatQualifiedQuery(t *testing.T) {
	t.Parallel()

	t.Run(
		"when using one parameter",
		func(t *testing.T) {
			t.Parallel()
			req := geocodingsearchv7.AddressRequest{
				Country: "Sweden",
			}
			res := geocodingsearchv7.FormatQualifiedQuery(req)
			assert.Equal(t, res, "country=Sweden")
		},
	)

	t.Run(
		"when using two parameters",
		func(t *testing.T) {
			t.Parallel()
			req := geocodingsearchv7.AddressRequest{
				Country: "Sweden",
				City:    "Stockholm",
			}
			res := geocodingsearchv7.FormatQualifiedQuery(req)
			assert.Equal(t, res, "country=Sweden;city=Stockholm")
		},
	)

	t.Run(
		"when using multiple parameters",
		func(t *testing.T) {
			t.Parallel()
			req := geocodingsearchv7.AddressRequest{
				Country:     "Sweden",
				City:        "Stockholm",
				Street:      "Regeringsgatan",
				HouseNumber: "65",
			}
			res := geocodingsearchv7.FormatQualifiedQuery(req)
			assert.Equal(t, res, "country=Sweden;city=Stockholm;street=Regeringsgatan;houseNumber=65")
		},
	)

	t.Run(
		"when using no parameters, then return empty qq",
		func(t *testing.T) {
			t.Parallel()
			req := geocodingsearchv7.AddressRequest{}
			res := geocodingsearchv7.FormatQualifiedQuery(req)
			assert.Equal(t, res, "")
		},
	)
}
