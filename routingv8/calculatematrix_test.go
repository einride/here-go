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

func TestCalculateMatrixBody_MarshalJSON(t *testing.T) {
	t.Parallel()
	for _, tt := range []struct {
		name string
		body routingv8.CalculateMatrixBody
		want string
	}{
		{
			name: "bare coordinates omit snap hints",
			body: routingv8.CalculateMatrixBody{
				Origins:      []*routingv8.GeoWaypoint{{Lat: 57.707752, Long: 11.949767}},
				Destinations: []*routingv8.GeoWaypoint{{Lat: 59.337492, Long: 18.063672}},
				RegionDefinition: routingv8.RegionDefinition{
					Type: routingv8.RegionTypeWorld,
				},
				TransportMode: routingv8.TransportModeTruck,
			},
			want: `{` +
				`"origins":[{"lat":57.707752,"lng":11.949767}],` +
				`"destinations":[{"lat":59.337492,"lng":18.063672}],` +
				`"regionDefinition":{"type":"world"},` +
				`"transportMode":"truck"` +
				`}`,
		},
		{
			name: "snap hints emit radius and radiusPenalty",
			body: routingv8.CalculateMatrixBody{
				Origins: []*routingv8.GeoWaypoint{
					{Lat: 57.707752, Long: 11.949767, Radius: 200, RadiusPenalty: 5000},
				},
				Destinations: []*routingv8.GeoWaypoint{
					{Lat: 59.337492, Long: 18.063672, Radius: 200, RadiusPenalty: 5000},
				},
				RegionDefinition: routingv8.RegionDefinition{
					Type: routingv8.RegionTypeWorld,
				},
				TransportMode: routingv8.TransportModeTruck,
			},
			want: `{` +
				`"origins":[{"lat":57.707752,"lng":11.949767,"radius":200,"radiusPenalty":5000}],` +
				`"destinations":[{"lat":59.337492,"lng":18.063672,"radius":200,"radiusPenalty":5000}],` +
				`"regionDefinition":{"type":"world"},` +
				`"transportMode":"truck"` +
				`}`,
		},
		{
			name: "snapRadius emits snapRadius",
			body: routingv8.CalculateMatrixBody{
				Origins: []*routingv8.GeoWaypoint{
					{Lat: 57.707752, Long: 11.949767, SnapRadius: 50},
				},
				Destinations: []*routingv8.GeoWaypoint{
					{Lat: 59.337492, Long: 18.063672, SnapRadius: 50},
				},
				RegionDefinition: routingv8.RegionDefinition{
					Type: routingv8.RegionTypeWorld,
				},
				TransportMode: routingv8.TransportModeTruck,
			},
			want: `{` +
				`"origins":[{"lat":57.707752,"lng":11.949767,"snapRadius":50}],` +
				`"destinations":[{"lat":59.337492,"lng":18.063672,"snapRadius":50}],` +
				`"regionDefinition":{"type":"world"},` +
				`"transportMode":"truck"` +
				`}`,
		},
		{
			name: "course hints emit course and minCourseDistance",
			body: routingv8.CalculateMatrixBody{
				Origins: []*routingv8.GeoWaypoint{
					{Lat: 57.707752, Long: 11.949767, Course: 90, MinCourseDistance: 500},
				},
				Destinations: []*routingv8.GeoWaypoint{
					{Lat: 59.337492, Long: 18.063672},
				},
				RegionDefinition: routingv8.RegionDefinition{
					Type: routingv8.RegionTypeWorld,
				},
				TransportMode: routingv8.TransportModeTruck,
			},
			want: `{` +
				`"origins":[{"lat":57.707752,"lng":11.949767,"course":90,"minCourseDistance":500}],` +
				`"destinations":[{"lat":59.337492,"lng":18.063672}],` +
				`"regionDefinition":{"type":"world"},` +
				`"transportMode":"truck"` +
				`}`,
		},
		{
			name: "sideOfStreetHint embeds match",
			body: routingv8.CalculateMatrixBody{
				Origins: []*routingv8.GeoWaypoint{
					{
						Lat:  52.511496,
						Long: 13.304140,
						SideOfStreetHint: &routingv8.SideOfStreetHint{
							Lat:   52.512149,
							Long:  13.304076,
							Match: routingv8.SideOfStreetMatchAlways,
						},
					},
				},
				Destinations: []*routingv8.GeoWaypoint{
					{Lat: 59.337492, Long: 18.063672},
				},
				RegionDefinition: routingv8.RegionDefinition{
					Type: routingv8.RegionTypeWorld,
				},
				TransportMode: routingv8.TransportModeTruck,
			},
			want: `{` +
				`"origins":[{"lat":52.511496,"lng":13.30414,` +
				`"sideOfStreetHint":{"lat":52.512149,"lng":13.304076,"match":"always"}}],` +
				`"destinations":[{"lat":59.337492,"lng":18.063672}],` +
				`"regionDefinition":{"type":"world"},` +
				`"transportMode":"truck"` +
				`}`,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := json.Marshal(&tt.body)
			assert.NilError(t, err)
			assert.Equal(t, tt.want, string(got))
		})
	}
}
