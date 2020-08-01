package routingv7

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type CalculateMatrixRequest struct {
	// StartWaypoints defining start points of the routes in the matrix.
	// The number of starts in M:N requests is limited to 15. In M:1 requests it is limited to 100.
	StartWaypoints []WaypointParameter
	// DestinationWaypoints defining destinations of the routes in the matrix.
	// The number of destinations in one request is limited to 100.
	// In case of pedestrian mode the number is limited by capacity of the request.
	DestinationWaypoints []WaypointParameter
	// The routing mode determines how the route is calculated.
	Mode RoutingMode
	// SummaryAttributes defines which attributes are included in the response as part of the data representation of
	// the matrix entries summaries. Defaults to cost factor.
	SummaryAttributes []MatrixRouteSummaryAttribute
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

func (r *CalculateMatrixRequest) Encode() string {
	var values url.Values
	for i, wp := range r.StartWaypoints {
		values.Add(fmt.Sprintf("start%d", i), wp.QueryString())
	}
	for i, wp := range r.DestinationWaypoints {
		values.Add(fmt.Sprintf("destination%d", i), wp.QueryString())
	}
	mode := r.Mode.String()
	if mode != "" {
		values.Add("mode", r.Mode.String())
	}
	for _, attr := range r.SummaryAttributes {
		values.Add("summaryAttributes", attr.String())
	}
	if r.TruckType != TruckTypeInvalid {
		values.Add("truckType", r.TruckType.String())
	}
	if r.TrailersCount > 0 {
		values.Add("trailersCount", strconv.Itoa(r.TrailersCount))
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

// CalculateMatrixResponse is used to provide results of a matrix calculation.
type CalculateMatrixResponse struct {
	// MetaInfo provides details about the request itself, such as the time at which it was processed, a request id,
	// or the map version on which the calculation was based.
	MetaInfo RouteResponseMetaInfo `json:"metaInfo,omitempty"`
	// MatrixEntries Entries are the entries in the matrix.
	// There is one entry for each start and each destination which provides the corresponding route.
	// The overall number of entries is therefore #starts * #destinations..
	MatrixEntries []RouteResponse `json:"matrixEntry,omitempty"`
}

// RouteMatrixEntry provides summary information for the route between waypoints indicated by indices of the start
// point and the destination. In cases of calculation failure the summary is replaced with a status field.
type RouteMatrixEntry struct {
	// StartIndex to identify the start point of the route in the list of starting positions.
	StartIndex int `json:"startIndex,omitempty"`
	// DestinationIndex to identify the destination point of the route in the list of destinations.
	DestinationIndex int `json:"destinationIndex,omitempty"`
	// Summary for this matrix entry. In the CalculateMatrix response, depending on summaryAttributes parameter, it may
	// include fields CostFactor, Distance and TravelTime.
	// The CalculateMatrix response does not provide other RouteSummaryType attributes.
	// The summary is included in the response only if route calculation succeeds.
	Summary MatrixRouteSummary `json:"summary,omitempty"`
	// Status of the route calculation.
	Status RouteStatus `json:"status,omitempty"`
}

// MatrixRouteSummary is used in calculate matrix responses.
type MatrixRouteSummary struct {
	// Distance indicates total travel distance for the route, in meters.
	Distance float64 `json:"distance,omitempty"`
	// Total travel time in seconds optionally considering traffic depending on the request parameters.
	TravelTime float64 `json:"travelTime,omitempty"`
	// CostFactor is an internal cost used for calculating the route.
	// This value is based on the objective function of the routing engine and related to the distance or time,
	// depending on the request settings (such as pedestrian versus car routes).
	// The value may include certain penalties and represents the overall quality of the route with respect to the
	// input parameters.
	// Computing the CostFactor value is computationally less expensive compared to calculating distance or time.
	// Especially for computing matrix routes and 1:n routes this value is preferred over distance or travel time,
	// as it allows computing responses in much less time. CostFactor is also the preferred metric for comparing
	// multiple routes such as for doing trip planning (traveling salesman problem) based on the response of a matrix
	// route calculation.
	// If time and distance is required for only a subset of a matrix route response, it is recommended to query for
	// distance and time separately, instead of requesting these values directly for a matrix request.
	CostFactor int `json:"costFactor,omitempty"`
	// RouteID is a unique identifier for the Get Route resource to get a corresponding route.
	// Consider RouteID as a semi-persistent reference to the route.
	// It may be incompatible between distinct map versions.
	RouteID string `json:"routeId,omitempty"`
}

// RouteStatus indicates the reason for a failed route calculation.
type RouteStatus string

const (
	// RouteStatusSuccess indicates a successful route calculation.
	RouteStatusSuccess = ""
	// RouteStatusRangeExceeded indicates that a route calculation was aborted due to the exceeded search range
	// specified in the request.
	RouteStatusRangeExceeded = "rangeExceeded"
	// RouteStatusFailed indicates that a route calculation failed due to invalid or non-reachable destination waypoint.
	RouteStatusFailed = "failed"
)

// CalculateMatrix returns a matrix of route summaries.
// The required parameters for this resource are mode and a set of start and destination waypoints.
// Other parameters can be left unspecified.
func (s *RouteService) CalculateMatrix(
	ctx context.Context,
	req *CalculateMatrixRequest,
) (*CalculateMatrixResponse, error) {
	params := req.Encode()
	u, err := s.URL.Parse("calculatematrix.json")
	if err != nil {
		return nil, err
	}
	r, err := s.client.NewRequest(ctx, u, http.MethodGet, params, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create get request: %v", err)
	}
	var resp struct {
		Response CalculateMatrixResponse `json:"response"`
	}
	if err = s.client.Do(r, &resp); err != nil {
		return nil, fmt.Errorf("unable to get routes: %v", err)
	}
	return &resp.Response, nil
}
