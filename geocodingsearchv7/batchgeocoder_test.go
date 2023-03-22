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

type BatchGeocodingMock struct {
	responseStatus int
	responseBody   geocodingsearchv7.BatchGeocoderResponse
}

func (c *BatchGeocodingMock) Do(_ *http.Request) (*http.Response, error) {
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

func TestBatchGeocodingService_BatchGeocode(t *testing.T) {
	t.Parallel()

	t.Run(
		"when both addresses and queries are nil, then return InvalidArgument",
		func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			httpClient := BatchGeocodingMock{responseBody: geocodingsearchv7.BatchGeocoderResponse{}, responseStatus: 200}
			routingClient := geocodingsearchv7.NewClient(&httpClient)

			_, err := routingClient.BatchGeocoding.BatchGeocoderUpload(ctx, &geocodingsearchv7.BatchGeocoderUploadRequest{
				Addresses: nil,
				Queries:   nil,
			})
			assert.ErrorContains(t, err, "InvalidArgument")
		},
	)
	t.Run(
		"when both addresses and queries are not set, then return InvalidArgument ",
		func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			httpClient := BatchGeocodingMock{responseBody: geocodingsearchv7.BatchGeocoderResponse{}, responseStatus: 200}
			routingClient := geocodingsearchv7.NewClient(&httpClient)

			_, err := routingClient.BatchGeocoding.BatchGeocoderUpload(ctx, &geocodingsearchv7.BatchGeocoderUploadRequest{
				Addresses: []*geocodingsearchv7.AddressRequest{},
				Queries:   []*geocodingsearchv7.QueryString{},
			})
			assert.ErrorContains(t, err, "InvalidArgument")
		},
	)
}

func TestBatchGeocodingService_BatchGeocodeStatus(t *testing.T) {
	t.Parallel()

	t.Run(
		"when requestID is empty, then return InvalidArgument",
		func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			httpClient := BatchGeocodingMock{responseBody: geocodingsearchv7.BatchGeocoderResponse{}, responseStatus: 200}
			routingClient := geocodingsearchv7.NewClient(&httpClient)

			_, err := routingClient.BatchGeocoding.BatchGeocoderStatus(ctx, &geocodingsearchv7.BatchGeocoderStatusRequest{
				RequestID: "",
			})
			assert.ErrorContains(t, err, "InvalidArgument")
		},
	)
}

func TestBatchGeocodingService_BatchGeocodeDownload(t *testing.T) {
	t.Parallel()

	t.Run(
		"when requestID is empty, then return InvalidArgument",
		func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			httpClient := BatchGeocodingMock{responseBody: geocodingsearchv7.BatchGeocoderResponse{}, responseStatus: 200}
			routingClient := geocodingsearchv7.NewClient(&httpClient)

			err := routingClient.BatchGeocoding.BatchGeocoderDownload(
				ctx,
				&geocodingsearchv7.BatchGeocoderDownloadRequest{RequestID: ""},
				nil,
			)
			assert.ErrorContains(t, err, "InvalidArgument")
		},
	)
}
