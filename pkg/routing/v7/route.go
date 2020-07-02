package routingv7

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// RouteService handles communication with the route related methods of the
// v7 HERE API.
type RouteService service

type CalculateRouteRequest struct {
	// List of waypoints that define a route. The first element marks the startpoint,
	// the last the endpoint. Waypoints in between are interpreted as via points.
	Waypoints []WaypointParameter
	// The routing mode determines how the route is calculated.
	Mode RoutingMode
}

// CalculateRoute calculates a route using a generic vehicle/pedestrian mode.
// e.g. car, truck, pedestrian, etc...
func (s *RouteService) CalculateRoute(
	ctx context.Context,
	req *CalculateRouteRequest,
) (*CalculateRouteResponse, error) {
	vals := url.Values{}
	for i, wp := range req.Waypoints {
		vals.Add(fmt.Sprintf("waypoint%d", i), wp.QueryString())
	}
	mode := req.Mode.String()
	if mode != "" {
		vals.Add("mode", req.Mode.String())
	}
	u, err := s.URL.Parse("calculateroute.json")
	if err != nil {
		return nil, err
	}
	r, err := s.client.NewRequest(ctx, u, http.MethodGet, vals.Encode(), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create get request: %v", err)
	}
	var resp struct {
		Response CalculateRouteResponse `json:"response"`
	}
	if err = s.client.Do(r, &resp); err != nil {
		return nil, fmt.Errorf("unable to get routes: %v", err)
	}
	return &resp.Response, nil
}
