package geocodingsearchv7

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// BatchGeocoderUpload allows batch forward geocoding of addresses.
// See https://developer.here.com/documentation/batch-geocoder/dev_guide/topics/introduction.html
// for details about other parameters.
func (s *BatchGeocodingService) BatchGeocoderUpload(
	ctx context.Context,
	req *BatchGeocoderUploadRequest,
) (_ *BatchGeocoderResponse, err error) {
	if req.Addresses != nil && req.Queries != nil {
		return nil, fmt.Errorf(
			"InvalidArgument, only one of Addresses or Queries can be used in the same request",
		)
	}
	if req.Addresses == nil && req.Queries == nil {
		return nil, fmt.Errorf("InvalidArgument, one of Addresses or Queries must be supplied")
	}
	u, err := s.URL.Parse("jobs")
	if err != nil {
		return nil, err
	}

	values := make(url.Values)
	values.Add("action", "run")
	values.Add("indelim", "|")
	values.Add("outdelim", "|")
	values.Add("outcols", "displayLatitude,displayLongitude,locationLabel,houseNumber,street,"+
		"district,city,postalCode,county,state,country")
	values.Add("outputCombined", "false")

	var body []byte
	if req.Addresses != nil {
		body = geoAddressesBody(req.Addresses)
	} else if req.Queries != nil {
		body = queryBody(req.Queries)
	}

	r, err := s.Client.NewRequest(ctx, u, http.MethodPost, values.Encode(), body)
	if err != nil {
		return nil, fmt.Errorf("unable to create post request: %v", err)
	}
	var resp BatchGeocoderResponse
	if err := s.Client.DoXML(r, &resp); err != nil {
		return nil, fmt.Errorf("DoXML failed: %v", err)
	}
	return &resp, nil
}

// BatchReverseGeocoderUpload allows batch reverse geocoding of geo-positions.
// See https://developer.here.com/documentation/batch-geocoder/dev_guide/topics/introduction.html
// for details about other parameters.
func (s *BatchGeocodingService) BatchReverseGeocoderUpload(
	ctx context.Context,
	req *BatchReverseGeocoderUploadRequest,
) (_ *BatchGeocoderResponse, err error) {
	if req.GeoPositions == nil {
		return nil, fmt.Errorf("InvalidArgument, geoPositions must be in the request")
	}
	u, err := s.URL.Parse("jobs")
	if err != nil {
		return nil, err
	}

	values := make(url.Values)
	values.Add("action", "run")
	values.Add("indelim", "|")
	values.Add("outdelim", "|")
	values.Add("outcols", "displayLatitude,displayLongitude,locationLabel,houseNumber,street,"+
		"district,city,postalCode,county,state,country")
	values.Add("outputCombined", "false")
	// Used to signal to Here that reverse geocoding should be performed
	values.Add("mode", "retrieveAddresses")

	body := geoPositionBody(req.GeoPositions)

	r, err := s.Client.NewRequest(ctx, u, http.MethodPost, values.Encode(), body)
	if err != nil {
		return nil, fmt.Errorf("unable to create post request: %v", err)
	}
	var resp BatchGeocoderResponse
	if err := s.Client.DoXML(r, &resp); err != nil {
		return nil, fmt.Errorf("DoXML failed: %v", err)
	}
	return &resp, nil
}

// BatchGeocoderStatus check status of batch geocoder job. The requestId was obtained in the BatchGeocoderUpload call
// that started the job.
// https://developer.here.com/documentation/examples/rest/batch_geocoding/batch-geocode-job-status
// for details about other parameters.
func (s *BatchGeocodingService) BatchGeocoderStatus(
	ctx context.Context,
	req *BatchGeocoderStatusRequest,
) (_ *BatchGeocoderResponse, err error) {
	if req.RequestID == "" {
		return nil, fmt.Errorf("InvalidArgument, requestID can not be empty")
	}
	u, err := s.URL.Parse(fmt.Sprintf("jobs/%s", req.RequestID))
	if err != nil {
		return nil, err
	}

	values := make(url.Values)
	values.Add("action", "status")

	r, err := s.Client.NewRequest(ctx, u, http.MethodGet, values.Encode(), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create get request: %v", err)
	}
	var resp BatchGeocoderResponse
	if err := s.Client.DoXML(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// BatchGeocoderDownload downloads the results of the completed batch geocoding job and returns the results as
// a zipped csv file.
// The requestId was obtained in the BatchGeocoderUpload call that started the job.
// Return an io.Writer with the zipped csv file.
// https://developer.here.com/documentation/examples/rest/batch_geocoding/download-geocoded-data
// for details about other parameters.
func (s *BatchGeocodingService) BatchGeocoderDownload(
	ctx context.Context,
	req *BatchGeocoderDownloadRequest,
	w io.Writer,
) error {
	if req.RequestID == "" {
		return fmt.Errorf("InvalidArgument, requestID can not be empty")
	}
	u, err := s.URL.Parse(fmt.Sprintf("jobs/%s/result", req.RequestID))
	if err != nil {
		return err
	}
	r, err := s.Client.NewRequest(ctx, u, http.MethodGet, "", nil)
	if err != nil {
		return fmt.Errorf("unable to create get request: %v", err)
	}
	if err := s.Client.DoXML(r, w); err != nil {
		return err
	}
	return nil
}

func geoPositionBody(p []*GeoWaypointRequest) []byte {
	var b []byte
	cols := "recId|prox"
	b = []byte(cols)
	for _, e := range p {
		b = append(b, []byte(fmt.Sprintf("\n%v|%v,%v", e.RecID, e.GeoPositions.Lat, e.GeoPositions.Long))...)
	}
	return b
}

func geoAddressesBody(addresses []*AddressRequest) []byte {
	var b []byte
	cols := "recId|street|houseNumber|district|city|postalCode|county|state|country"
	b = []byte(cols)
	for _, a := range addresses {
		address := fmt.Sprintf(
			"%s|%s|%s|%s|%s|%s|%s|%s",
			a.Street,
			a.HouseNumber,
			a.District,
			a.City,
			a.PostalCode,
			a.County,
			a.State,
			a.Country,
		)
		b = append(b, []byte(fmt.Sprintf("\n%v|%s", a.RecID, address))...)
	}
	return b
}

func queryBody(queries []*QueryString) []byte {
	var b []byte
	cols := "recId|searchText|country"
	b = []byte(cols)
	for _, a := range queries {
		b = append(b, []byte(fmt.Sprintf("\n%v|%s|%s", a.RecID, a.Query, a.Country))...)
	}
	return b
}
