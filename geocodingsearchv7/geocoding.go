package geocodingsearchv7

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// Geocoding allows forward geocoding of addresses and geo-positions.
// See https://developer.here.com/documentation/geocoding-search-api/dev_guide/topics/quick-start.html#send-a-request
// for details about other parameters.
func (s *GeocodingService) Geocoding(
	ctx context.Context,
	req *GeocodingRequest,
) (_ *GeocodingResponse, err error) {
	u, err := s.URL.Parse("geocode")
	if err != nil {
		return nil, err
	}

	if req.Q == nil && req.Address == nil {
		return nil, fmt.Errorf("InvalidArgument, either Queries or QQ must be provided")
	}

	values := make(url.Values)
	if req.GeoPosition != nil {
		values.Add("at", fmt.Sprintf("%v,%v", req.GeoPosition.Lat, req.GeoPosition.Long))
	}
	if req.Q != nil {
		values.Add("q", *req.Q)
	}
	if req.Address != nil {
		values.Add("qq", FormatQualifiedQuery(*req.Address))
	}
	if req.In != nil {
		values.Add("in", *req.In)
	}

	r, err := s.Client.NewRequest(ctx, u, http.MethodGet, values.Encode(), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create get request: %v", err)
	}
	var resp GeocodingResponse
	if err := s.Client.Do(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// FormatQualifiedQuery takes an address and formats it into a qualified query.
// The resulting string will look like <key1>=<value1>;<key2>=<value2>;...
// where the keys are all fields of the address object and the values are their values.
func FormatQualifiedQuery(address AddressRequest) string {
	var qq string
	if address.Country != "" {
		qq = addSubParam(qq, "country", address.Country)
	}
	if address.State != "" {
		qq = addSubParam(qq, "state", address.State)
	}
	if address.County != "" {
		qq = addSubParam(qq, "county", address.County)
	}
	if address.City != "" {
		qq = addSubParam(qq, "city", address.City)
	}
	if address.District != "" {
		qq = addSubParam(qq, "district", address.District)
	}
	if address.Street != "" {
		qq = addSubParam(qq, "street", address.Street)
	}
	if address.HouseNumber != "" {
		qq = addSubParam(qq, "houseNumber", address.HouseNumber)
	}
	if address.PostalCode != "" {
		qq = addSubParam(qq, "postalCode", address.PostalCode)
	}
	return qq
}

func addSubParam(query string, paramName string, value string) string {
	if query != "" {
		query = fmt.Sprintf("%s;", query)
	}
	query = fmt.Sprintf("%s%s=%s", query, paramName, value)
	return query
}
