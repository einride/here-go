package routingv8_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"go.einride.tech/here/routingv8"
	"gotest.tools/v3/assert"
)

type ClientMock struct {
	responseStatus int
	responseBody   routingv8.CalculateMatrixResponse
}

func (c *ClientMock) Do(_ *http.Request) (*http.Response, error) {
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

func TestMatrixService_CalculateMatrix(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	exp := routingv8.CalculateMatrixResponse{
		MatrixID: "123",
		Matrix: routingv8.MatrixResponse{
			NumOrigins:      1,
			NumDestinations: 1,
			TravelTimes:     []int32{},
			Distances:       []int32{1},
			ErrorCodes:      routingv8.ErrorCodes{routingv8.ErrorCodeSuccess},
		},
		RegionDefinition: routingv8.RegionDefinition{
			Type: routingv8.RegionTypeWorld,
		},
	}
	httpClient := ClientMock{responseBody: exp, responseStatus: 200}
	routingClient := routingv8.NewClient(&httpClient)
	// Einride Gothenburg.
	origins := []*routingv8.GeoWaypoint{
		{
			Lat:  57.707752,
			Long: 11.949767,
		},
	}
	// Einride Stockholm.
	destinations := []*routingv8.GeoWaypoint{
		{
			Lat:  59.337492,
			Long: 18.063672,
		},
	}
	got, err := routingClient.Matrix.CalculateMatrix(ctx, &routingv8.CalculateMatrixRequest{
		Async: false,
		Body: &routingv8.CalculateMatrixBody{
			Origins:      origins,
			Destinations: destinations,
			RegionDefinition: routingv8.RegionDefinition{
				Type: routingv8.RegionTypeWorld,
			},
			Profile: routingv8.ProfileTruckFast,
			MatrixAttributes: &routingv8.MatrixAttributes{
				routingv8.MatrixAttributeDistances,
				routingv8.MatrixAttributeTravelTimes,
			},
		},
	})
	assert.NilError(t, err)
	assert.DeepEqual(t, &exp, got)
}
