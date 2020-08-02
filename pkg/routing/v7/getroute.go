package routingv7

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type GetRouteRequest struct {
	// RouteID is the route identifier for which the detailed route information is being requested.
	RouteID string
	// List of waypoints that define a route. The first element marks the startpoint,
	// the last the endpoint. Waypoints in between are interpreted as via points.
	Waypoints []WaypointParameter
	// The routing mode determines how the route is calculated.
	Mode RoutingMode
	// Representation defines which elements are included in the response
	// as part of the data representation of the route.
	Representation RouteRepresentationMode
}

// GetRouteResponse contains response data, structured to match a particular request for the GetRoute operation.
type GetRouteResponse struct {
	// MetaInfo provides details about the request itself, such as the time at which it was processed, a request id,
	// or the map version on which the calculation was based.
	MetaInfo RouteResponseMetaInfo `json:"metaInfo,omitempty"`
	// Contains the calculated path across a navigable link network, as specified in the request.
	Route Route `json:"route,omitempty"`
}

func (r *GetRouteRequest) QueryString() string {
	values := make(url.Values)
	for i, wp := range r.Waypoints {
		values.Add(fmt.Sprintf("waypoint%d", i), wp.QueryString())
	}
	if mode := r.Mode.String(); mode != "" {
		values.Add("mode", mode)
	}
	if representation := r.Representation.String(); representation != "" {
		values.Add("representation", representation)
	}
	return values.Encode()
}

// GetRoute requests a previously calculated route by providing a route ID.
//
// As currently calculation of RouteId for Public Transport is not possible, GetRoute cannot be used for
// Public Transport.
func (s *RouteService) GetRoute(
	ctx context.Context,
	req *GetRouteRequest,
) (*GetRouteResponse, error) {
	u, err := s.URL.Parse("getroute.json")
	if err != nil {
		return nil, err
	}
	r, err := s.client.NewRequest(ctx, u, http.MethodGet, req.QueryString(), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create get request: %v", err)
	}
	var resp struct {
		Response GetRouteResponse `json:"response"`
	}
	if err = s.client.Do(r, &resp); err != nil {
		return nil, fmt.Errorf("unable to get routes: %v", err)
	}
	return &resp.Response, nil
}
