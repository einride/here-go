package routingv7

import (
	"fmt"
	"strconv"
	"strings"
)

type WaypointParameter interface {
	QueryString() string
}

type WaypointTypeEnum int

const (
	WaypointTypInvalid WaypointTypeEnum = iota
	WaypointTypeStopOver
	WaypointTypePassthrough
)

type WaypointType struct {
	// 180 degree turns are allowed for WaypointTypeStopOver but not for WaypointTypePassthrough.
	// Waypoints defined through a drag-n-drop action should be marked as pass-through.
	// PassThrough waypoints will not appear in the list of maneuvers.
	Type WaypointTypeEnum
	// Stopover delay in seconds. Impacts time-aware calculations. Ignored for WaypointTypePassthrough.
	StopOverDuration int
}

func (t WaypointType) String() string {
	var out string
	switch t.Type {
	case WaypointTypePassthrough:
		out = "passThrough!"
	case WaypointTypeStopOver:
		out = "stopOver"
		if t.StopOverDuration > 0 {
			out += ","
			out += strconv.Itoa(t.StopOverDuration)
		}
		out += "!"
	}
	return out
}

type GeoWaypoint struct {
	Lat          float64
	Long         float64
	WaypointType WaypointType
	// In meters
	Altitude int32
	// Matching Links are selected within the specified TransitRadius.
	// For example to drive past a city without necessarily going into the city center you can
	// specify the coordinates of the center and a TransitRadius of 5000m.
	// In meters
	TransitRadius int
	// Custom label identifying this waypoint.
	UserLabel string
	// Heading (at the startpoint), in degrees starting at true north and continuing clockwise around the compass.
	// North is 0 degrees, East is 90 degrees, South is 180 degrees, and West is 270 degrees.
	// MUST be in range [0,360]
	Heading float32
}

func (w *GeoWaypoint) QueryString() string {
	builder := strings.Builder{}
	builder.WriteString("geo!")
	if wt := w.WaypointType.String(); wt != "" {
		builder.WriteString(wt)
	}
	builder.WriteString(strconv.FormatFloat(w.Lat, 'f', 8, 64))
	builder.WriteString(",")
	builder.WriteString(strconv.FormatFloat(w.Long, 'f', 8, 64))
	if w.TransitRadius > 0 {
		builder.WriteString(";")
		builder.WriteString(strconv.Itoa(w.TransitRadius))
	}
	if w.UserLabel != "" {
		builder.WriteString(";")
		builder.WriteString(w.UserLabel)
	}
	if w.Heading > 0 {
		builder.WriteString(";")
		builder.WriteString(fmt.Sprintf("%f", w.Heading))
	}
	return builder.String()
}

type RoutingMode struct {
	// Type provides identifiers for different optimizations which can be applied during the route calculation.
	// Selecting the routing type affects which constraints, speed attributes and weights
	// are taken into account during route calculation.
	// Defaults to RouteTypeFastest.
	Type RouteType
	// Depending on the transport mode special constraints, speed attributes and weights are taken into account during
	// route calculation.
	// Defaults to TransportModeCar.
	TransportMode TransportModeType
	// Specify whether to optimize a route for traffic.
	// Note: Difference between traffic disabled and enabled affects only the calculation of the route.
	// Traffic time of the route will still be calculated for all routes using the same rules as for traffic:enabled.
	// Optional.
	TrafficMode TrafficModeType
	// TODO: Implement RouteFeature
	// https://developer.here.com/cn/documentation/routing/topics/resource-param-type-routing-mode.html#type-route-feature
}

func (m RoutingMode) String() string {
	var out string
	out += m.Type.String()
	out += ";"
	out += m.TransportMode.String()
	out += ";"
	if m.TrafficMode != TrafficModeTypeInvalid {
		out += "traffic:"
		out += m.TrafficMode.String()
		out += ";"
	}
	return out
}

type RouteType int

const (
	// Route calculation from start to destination optimized by travel time.
	// Depending on the traffic mode provided by the request, travel time is determined with or without
	// traffic information.
	// In some cases, the route returned by the fastest mode may not be the route with the shortest possible
	// travel time.
	// For example, the routing service may favor a route that remains on a highway, even if a shorter travel
	// time can be achieved by taking a detour or shortcut through a side road.
	RouteTypeFastest RouteType = iota
	// Route calculation from start to destination disregarding any speed information. In this mode, the distance of the
	// route is minimized, while keeping the route sensible.
	// This includes, for example, penalizing turns. Because of that, the resulting route will not necessarily
	// be the one with minimal distance.
	RouteTypeShortest
	// Route calculation from start to destination optimizing based on combination of travel time and distance.
	RouteTypeBalanced
)

func (t RouteType) String() string {
	var out string
	switch t {
	case RouteTypeFastest:
		out = "fastest"
	case RouteTypeShortest:
		out = "shortest"
	case RouteTypeBalanced:
		out = "balanced"
	}
	return out
}

type TransportModeType int

const (
	// Route calculation for cars.
	TransportModeCar TransportModeType = iota
	// Route calculation for HOV (high-occupancy vehicle) cars.
	TransportModeCarHOV
	// Route calculation for a pedestrian. As one effect, maneuvers will be optimized for walking,
	// that is segments will consider actions relevant for pedestrians and maneuver instructions will
	// contain texts suitable for a walking person. This mode disregards any traffic information.
	TransportModePedestrian
	// Route calculation using public transport lines and walking parts to get to stations.
	// It is based on static map data, so the results are not aligned with officially published
	// timetable.
	TransportModePublicTransport
	// Route calculation using public transport lines and walking parts to get to stations.
	// This mode uses additional officially published timetable information to provide most precise
	// routes and times. In case the timetable data is unavailable, the service will use estimated
	// results based on static map data (same as from publicTransport mode).
	TransportModePublicTransportTimeTable
	// Route calculation for trucks. This mode considers truck limitations on links and uses different
	// speed assumptions when calculating the route.
	TransportModeTruck
	// Route calculation for bicycles. This mode uses the pedestrian road network, but uses different
	// speeds based on each road's suitability for cycling. Pedestrian roads that are also open for cars
	// in the travel direction are considered open for cycling, as are pedestrian roads located in parks.
	// These roads use full bicycle speed for routing. Other pedestrian roads, including travelling the wrong
	// way down one-way streets, are considered at the pedestrian walking speed, as it is assumed the bicycle
	// must be pushed.
	TransportModeBicycle
)

func (t TransportModeType) String() string {
	var out string
	switch t {
	case TransportModeCar:
		out = "car"
	case TransportModePedestrian:
		out = "pedestrian"
	case TransportModeCarHOV:
		out = "carHOV"
	case TransportModeBicycle:
		out = "bicycle"
	case TransportModePublicTransport:
		out = "publicTransport"
	case TransportModePublicTransportTimeTable:
		out = "publicTransportTimeTable"
	case TransportModeTruck:
		out = "truck"
	}
	return out
}

type TrafficModeType int

const (
	TrafficModeTypeInvalid TrafficModeType = iota
	// No departure time provided: This behavior is deprecated and will return error in the future.
	// 	* Static time based restrictions: Ignored
	// 	* Real-time traffic closures: Valid for entire length of route.
	// 	* Real-time traffic flow events: Speed at calculation time used for entire length of route.
	//
	// Departure time provided:
	// 	* Static time based restrictions: Avoided if road would be traversed within validity period of the restriction.
	// 	* Real-time traffic closures: Avoided if road would be traversed within validity period of the incident.
	// 	* Real-time traffic flow events: Speed used if road would be traversed within validity period of the flow event.
	TrafficModeTypeEnabled
	// No departure time provided:
	// 	* Static time based restrictions: Ignored
	// 	* Real-time traffic closures: Ignored.
	// 	* Real-time traffic flow speed: Ignored.
	//
	// Departure time provided:
	// 	* Static time based restrictions: Avoided if road would be traversed within validity period of the restriction.
	// 	* Real-time traffic closures: Valid for entire length of route.
	// 	* Real-time traffic flow speed: Ignored.
	TrafficModeTypeDisabled
	// Let the service automatically apply traffic related constraints that are suitable for the selected routing type,
	// transport mode, and departure time. Also user entitlements will be taken into consideration.
	TrafficModeTypeDefault
)

func (t TrafficModeType) String() string {
	var out string
	switch t {
	case TrafficModeTypeEnabled:
		out = "enabled"
	case TrafficModeTypeDisabled:
		out = "disabled"
	case TrafficModeTypeDefault:
		out = "default"
	}
	return out
}

type TruckType int

const (
	TruckTypeInvalid = iota
	TruckTypeTruck
	TruckTypeTractorTruck
)

func (t TruckType) String() string {
	var out string
	switch t {
	case TruckTypeTruck:
		out = "truck"
	case TruckTypeTractorTruck:
		out = "tractorTruck"
	}
	return out
}

// RouteRepresentationMode defines which parts of the route are returned by services for standard use cases.
type RouteRepresentationMode int

const (
	// RouteRepresentationModeOverview only returns the Route and the RouteSummary object.
	RouteRepresentationModeOverview RouteRepresentationMode = iota
	// RouteRepresentationModeDisplay that allows to show the route with all maneuvers.
	// Links will not be included in the response.
	RouteRepresentationModeDisplay
	// RouteRepresentationModeDragNDrop to be used during drag and drop (re-routing) actions.
	// The response will only contain the shape of the route restricted to the view bounds provided in the
	// representation options..
	RouteRepresentationModeDragNDrop
	// RouteRepresentationModeNavigation to provide all information necessary to support a navigation device.
	// This mode activates the most extensive data response as all link information will be included in the
	// response to allow a detailed display while navigating.
	// RouteId will not be calculated in this mode however, unless it is additionally requested..
	RouteRepresentationModeNavigation
	// RouteRepresentationModeLinkPaging that will be used when incrementally loading links
	// while navigating along the route.
	// The response will be limited to link information..
	RouteRepresentationModeLinkPaging
	// RouteRepresentationModeTurnByTurn mode to provide all information necessary to support turn by turn.
	// This mode activates all data needed for navigation excluding any redundancies.
	// RouteId will not be calculated in this mode however, unless it is additionally requested..
	RouteRepresentationModeTurnByTurn
)

func (r RouteRepresentationMode) String() string {
	switch r {
	case RouteRepresentationModeOverview:
		return "overview"
	case RouteRepresentationModeDisplay:
		return "display"
	case RouteRepresentationModeDragNDrop:
		return "dragNDrop"
	case RouteRepresentationModeNavigation:
		return "navigation"
	case RouteRepresentationModeLinkPaging:
		return "linkPaging"
	case RouteRepresentationModeTurnByTurn:
		return "turnByTurn"
	default:
		return ""
	}
}

// MatrixRouteSummaryAttribute defines which attributes are included in the response as part of the data
// representation of the matrix entries summaries. Defaults to cost factor.
type MatrixRouteSummaryAttribute int

const (
	// Indicates whether the travel time information should be provided in summary entries.
	MatrixRouteSummaryAttributeTravelTime MatrixRouteSummaryAttribute = iota
	// Indicates whether the CostFactor information should be returned in summary entries.
	MatrixRouteSummaryAttributeCostFactor
	// Indicates whether distance information should be returned in summary entries.
	MatrixRouteSummaryAttributeDistance
	// Indicates whether RouteId shall be calculated and provided in summary entries.
	MatrixRouteSummaryAttributeRouteID
)

func (m MatrixRouteSummaryAttribute) String() string {
	switch m {
	case MatrixRouteSummaryAttributeTravelTime:
		return "tt"
	case MatrixRouteSummaryAttributeCostFactor:
		return "cf"
	case MatrixRouteSummaryAttributeDistance:
		return "di"
	case MatrixRouteSummaryAttributeRouteID:
		return "ri"
	default:
		return ""
	}
}
