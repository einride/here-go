package geocodingsearchv7

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// ReverseGeocoding allows reverse geocode from geo-position to address.
// See https://developer.here.com/documentation/geocoding-search-api/api-reference-swagger.html
// for more details.
func (s *ReverseGeocodingService) ReverseGeocoding(
	ctx context.Context,
	req *ReverseGeocodingRequest,
) (*ReverseGeocodingResponse, error) {
	u, err := s.URL.Parse("revgeocode")
	if err != nil {
		return nil, err
	}

	if req.GeoPosition == nil {
		return nil, fmt.Errorf("InvalidArgument, GeoPosition must be provided")
	}

	values := make(url.Values)
	if req.GeoPosition != nil {
		values.Add("at", fmt.Sprintf("%v,%v", req.GeoPosition.Lat, req.GeoPosition.Long))
	}
	if req.In != nil {
		values.Add("in", *req.In)
	}

	r, err := s.Client.NewRequest(ctx, u, http.MethodGet, values.Encode(), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create get request: %v", err)
	}
	var resp ReverseGeocodingResponse
	if err := s.Client.Do(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
