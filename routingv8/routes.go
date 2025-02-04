package routingv8

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Routes returns all possible routes between origin and destination.
// See https://developer.here.com/documentation/routing-api/dev_guide/topics/send-request.html#send-a-request
// for details about other parameters.
func (s *RoutingService) Routes(
	ctx context.Context,
	req *RoutesRequest,
) (_ *RoutesResponse, err error) {
	tm := req.TransportMode.String()
	if tm == invalid || tm == unspecified {
		return nil, fmt.Errorf("invalid transportmode")
	}

	u, err := s.URL.Parse("routes")
	if err != nil {
		return nil, err
	}

	values := make(url.Values)
	returns := make([]string, 0, len(req.Return))
	for _, attribute := range req.Return {
		returns = append(returns, string(attribute))
	}
	if len(returns) == 0 {
		returns = append(returns, string(SummaryReturnAttribute))
	}
	values.Add("return", strings.Join(returns, ","))
	if req.DepartureTime != "" {
		values.Add("departureTime", req.DepartureTime)
	}
	values.Add("transportMode", tm)
	values.Add("origin", fmt.Sprintf("%v,%v", req.Origin.Lat, req.Origin.Long))
	values.Add("destination", fmt.Sprintf("%v,%v", req.Destination.Lat, req.Destination.Long))
	if len(req.Spans) > 0 {
		if !returnContains(req.Return, PolylineReturnAttribute) {
			return nil, errors.New("spans parameter also requires that the polyline option is set in the return parameter")
		}
		spanStrings := make([]string, 0, len(req.Spans))
		for _, span := range req.Spans {
			spanStrings = append(spanStrings, string(span))
		}
		values.Add("spans", strings.Join(spanStrings, ","))
	}
	if req.AvoidAreas != nil {
		areas := make([]string, 0, len(req.AvoidAreas))
		for _, area := range req.AvoidAreas {
			a := area.String()
			if a == invalid {
				return nil, fmt.Errorf("invalid avoid area")
			}
			if a != unspecified {
				areas = append(areas, a)
			}
		}
		values.Add("avoid[features]", strings.Join(areas, ","))
	}
	if req.Vehicle != nil {
		addVehicleParameters(values, req.Vehicle)
	}

	r, err := s.Client.NewRequest(ctx, u, http.MethodGet, values.Encode(), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create get request: %v", err)
	}
	var resp RoutesResponse
	if err := s.Client.Do(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func addVehicleParameters(values url.Values, vehicle *Vehicle) {
	if vehicle.GrossWeight != 0 {
		values.Add("vehicle[grossWeight]", strconv.Itoa(vehicle.GrossWeight))
	}
	if vehicle.TrailerCount != 0 {
		values.Add("vehicle[trailerCount]", strconv.Itoa(vehicle.TrailerCount))
	}
	if vehicle.AxleCount != 0 {
		values.Add("vehicle[axleCount]", strconv.Itoa(vehicle.AxleCount))
	}
	if vehicle.Height != 0 {
		values.Add("vehicle[height]", strconv.Itoa(vehicle.Height))
	}
	if vehicle.Width != 0 {
		values.Add("vehicle[width]", strconv.Itoa(vehicle.Width))
	}
	if vehicle.Length != 0 {
		values.Add("vehicle[length]", strconv.Itoa(vehicle.Length))
	}
	if vehicle.Type != "" {
		values.Add("vehicle[type]", vehicle.Type.String())
	}
}

// RouteImport returns a route from a sequence of trace points.
// See https://www.here.com/docs/bundle/routing-api-developer-guide-v8/page/concepts/route-import.html
// and https://www.here.com/docs/bundle/routing-api-v8-api-reference/page/index.html#tag/Routing/operation/importRoute
// for details.
func (s *RoutingService) RouteImport(
	ctx context.Context,
	req *RouteImportRequest,
) (_ *RoutesResponse, err error) {
	tm := req.TransportMode.String()
	if tm == invalid || tm == unspecified {
		return nil, fmt.Errorf("invalid transportmode")
	}

	if len(req.Trace) < 2 {
		return nil, fmt.Errorf("trace parameter must contain at least 2 waypoints")
	}

	u, err := s.URL.Parse("import")
	if err != nil {
		return nil, err
	}

	values := make(url.Values)
	returns := make([]string, 0, len(req.Return))
	for _, attribute := range req.Return {
		returns = append(returns, string(attribute))
	}
	if len(returns) == 0 {
		returns = append(returns, string(SummaryReturnAttribute))
	}
	values.Add("return", strings.Join(returns, ","))
	if req.DepartureTime != "" {
		values.Add("departureTime", req.DepartureTime)
	}
	values.Add("transportMode", tm)
	if len(req.Spans) > 0 {
		if !returnContains(req.Return, PolylineReturnAttribute) {
			return nil, errors.New("spans parameter also requires that the polyline option is set in the return parameter")
		}
		spanStrings := make([]string, 0, len(req.Spans))
		for _, span := range req.Spans {
			spanStrings = append(spanStrings, string(span))
		}
		values.Add("spans", strings.Join(spanStrings, ","))
	}
	if req.Vehicle != nil {
		addVehicleParameters(values, req.Vehicle)
	}

	bytes, err := json.Marshal(&RouteImportRequestBody{
		Trace: req.Trace,
	})
	if err != nil {
		return nil, err
	}

	r, err := s.Client.NewRequest(ctx, u, http.MethodPost, values.Encode(), bytes)
	if err != nil {
		return nil, fmt.Errorf("unable to create get request: %v", err)
	}
	var resp RoutesResponse
	if err := s.Client.Do(r, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func returnContains(requested []ReturnAttribute, needle ReturnAttribute) bool {
	for _, attr := range requested {
		if attr == needle {
			return true
		}
	}
	return false
}
