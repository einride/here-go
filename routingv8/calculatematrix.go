package routingv8

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func (c *CalculateMatrixRequest) QueryString() string {
	values := make(url.Values)
	values.Add("async", c.Async.String())
	return values.Encode()
}

// CalculateMatrix returns a matrix of route summaries.
// The required parameters for this resource are a region definition and a set of start and destination waypoints.
// See https://developer.here.com/documentation/matrix-routing-api/8.6.0/dev_guide/topics/get-started/send-request.html
// for details about other parameters.
func (s *MatrixService) CalculateMatrix(
	ctx context.Context,
	req *CalculateMatrixRequest,
) (_ *CalculateMatrixResponse, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("calculate matrix: %v", err)
		}
	}()
	u, err := s.URL.Parse("matrix")
	if err != nil {
		return nil, err
	}
	bytes, err := json.Marshal(req.Body)
	if err != nil {
		return nil, err
	}
	r, err := s.Client.NewRequest(ctx, u, http.MethodPost, req.QueryString(), bytes)
	if err != nil {
		return nil, fmt.Errorf("unable to create post request: %v", err)
	}
	var resp CalculateMatrixResponse
	if err := s.Client.Do(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
