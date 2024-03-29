package routingv7

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type CalculateRouteRequest struct {
	// List of waypoints that define a route. The first element marks the startpoint,
	// the last the endpoint. Waypoints in between are interpreted as via points.
	Waypoints []WaypointParameter
	// The routing mode determines how the route is calculated.
	Mode RoutingMode
	// Truck routing only, specifies the vehicle type. Defaults to truck.
	//
	// Note: Relevant for restrictions that apply exclusively to tractors with
	// semi-trailers (observed in North America). Such restrictions are taken into account
	// only in case of the truckType set to tractorTruck and the trailers count greater
	// than 0 (for example &truckType=tractorTruck&trailersCount=1). The truck type is irrelevant
	// in case of restrictions common for all types of trucks.
	TruckType TruckType
	// Truck routing only, specifies number of trailers pulled by the vehicle.
	// The provided value must be between 0 and 4. Defaults to 0.
	TrailersCount int
	// Truck routing only, specifies number of axles on the vehicle.
	// The provided value must be between 2 and 254. Defaults to 2.
	AxleCount int
	// Truck routing only, vehicle weight including trailers and shipped goods, in tons.
	// The provided value must be between 0 and 1000.
	LimitedWeight float64
	// Truck routing only, vehicle weight per axle in tons.
	// The provided value must be between 0 and 1000.
	WeightPerAxle float64
	// Truck routing only, vehicle height in meters.
	// The provided value must be between 0 and 50.
	Height float64
	// Truck routing only, vehicle width in meters.
	// The provided value must be between 0 and 50.
	Width float64
	// Truck routing only, vehicle length in meters.
	// The provided value must be between 0 and 300.
	Length float64
}

func (r *CalculateRouteRequest) QueryString() string {
	values := make(url.Values)
	for i, wp := range r.Waypoints {
		values.Add(fmt.Sprintf("waypoint%d", i), wp.QueryString())
	}
	if r.Mode.String() != "" {
		values.Add("mode", r.Mode.String())
	}
	if r.TruckType != TruckTypeInvalid {
		values.Add("truckType", r.TruckType.String())
	}
	if r.TrailersCount > 0 {
		values.Add("trailersCount", strconv.Itoa(r.TrailersCount))
	}
	if r.AxleCount > 0 {
		values.Add("axleCount", strconv.Itoa(r.AxleCount))
	}
	if r.LimitedWeight > 0 {
		values.Add("limitedWeight", strconv.FormatFloat(r.LimitedWeight, 'f', -1, 64))
	}
	if r.WeightPerAxle > 0 {
		values.Add("weightPerAxle", strconv.FormatFloat(r.WeightPerAxle, 'f', -1, 64))
	}
	if r.Height > 0 {
		values.Add("height", strconv.FormatFloat(r.Height, 'f', -1, 64))
	}
	if r.Width > 0 {
		values.Add("width", strconv.FormatFloat(r.Width, 'f', -1, 64))
	}
	if r.Length > 0 {
		values.Add("length", strconv.FormatFloat(r.Length, 'f', -1, 64))
	}
	return values.Encode()
}

// CalculateRouteResponse contains response data, structured to match a particular request for the CalculateRoute
// operation.
//
// The Route element may appear more than once in the response if multiple routes are available between the start and
// end points.
//
// Refer to CalculateRouteRequest, described in Calculate Route for details on associated request elements.
type CalculateRouteResponse struct {
	// MetaInfo provides details about the request itself, such as the time at which it was processed, a request id, or
	// the map version on which the calculation was based.
	MetaInfo RouteResponseMetaInfo `json:"metaInfo,omitempty"`
	// Routes contains the calculated path across a navigable link network, as specified in the request.
	// Routes contain navigation instructions for a single trip as: waypoints (fixed locations) and route legs
	// (sections of the route between waypoints). Each response may also include information about the route itself,
	// such as its overall shape, map location, or a summary description.
	Routes []Route `json:"route,omitempty"`
}

// CalculateRoute calculates a route using a generic vehicle/pedestrian mode.
// e.g. car, truck, pedestrian, etc...
func (s *RouteService) CalculateRoute(
	ctx context.Context,
	req *CalculateRouteRequest,
) (*CalculateRouteResponse, error) {
	u, err := s.URL.Parse("calculateroute.json")
	if err != nil {
		return nil, err
	}
	r, err := s.client.NewRequest(ctx, u, http.MethodGet, req.QueryString(), nil)
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
