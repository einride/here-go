package here

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// RoutingService handles communication with the routing related methods of the
// HERE API.
type RoutingService service

type RoutingRequest struct {
	// Mode of transport to be used for the calculation of the route.
	TransportMode TransportMode
	// A location defining origin of the trip.
	Origin Waypoint
	// A location defining destination of the trip.
	Destination Waypoint
	// A location defining a via waypoint.
	// A via waypoint is a location between origin and destination.
	// The route will do a stop at the via waypoint.
	Via *Waypoint
}

// Get calculates a route using a generic vehicle/pedestrian mode, e.g. car, truck, pedestrian, etc...
func (s *RoutingService) Get(ctx context.Context, req RoutingRequest) ([]byte, error) {
	vals := url.Values{}
	vals.Add("transportMode", req.TransportMode.String())
	vals.Add("origin", req.Origin.String())
	vals.Add("destination", req.Destination.String())
	if req.Via != nil {
		vals.Add("via", req.Via.String())
	}
	return s.getWithQueryAndPath(ctx, "routes", vals.Encode())
}

func (s *RoutingService) getWithQueryAndPath(
	ctx context.Context,
	path string,
	query string,
) ([]byte, error) {
	u, err := s.URL.Parse(path)
	if err != nil {
		return nil, err
	}
	req, err := s.client.NewRequest(ctx, u, http.MethodGet, query, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create get request: %v", err)
	}
	var buf bytes.Buffer
	err = s.client.Do(req, &buf)
	if err != nil {
		return nil, fmt.Errorf("unable to get routes: %v", err)
	}
	return buf.Bytes(), nil
}
