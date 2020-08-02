package routingv7

import "encoding/json"

type Route struct {
	// Permanent unique identifier of the route, generated based on route links.
	// Can be used to reproduce any previously calculated route.
	ID string `json:"routeId,omitempty"`

	// Overall route distance and time summary.
	Summary RouteSummary `json:"summary,omitempty"`

	// Legs is a partition of the route into legs between the different waypoints.
	Legs []RouteLeg `json:"leg,omitempty"`
}

type RouteSummary struct {
	// Indicates total travel distance for the route, in meters.
	DistanceMeters float64 `json:"distance,omitempty"`
	// Contains the travel time estimate for this element, considering traffic
	// and transport mode. Based on the TrafficSpeed. The service may also account for additional
	// time penalties, so this may be greater than the element length divided by the TrafficSpeed.
	TrafficTime Duration `json:"traffictime,omitempty"`
	// Contains the travel time estimate for this element, considering transport mode but
	// not traffic conditions. Based on the BaseSpeed. The service may also account for additional
	// time penalties, therefore this may be greater than the element length divided by the BaseSpeed.
	BaseTime Duration `json:"baseTime,omitempty"`
	// Special link characteristics (like ferry or motorway usage) which are covered by the route.
	Flags []RouteLinkFlag `json:"flags,omitempty"`
	// Total travel time optionally considering traffic depending on the request parameters.
	TravelTime Duration `json:"travelTime,omitempty"`
	// Textual description of route summary. Supported languages are US and GB English, French, German,
	// Italian, Polish, Portuguese, Spanish and Traditional Chinese. Element is not provided in case of
	// other languages.
	Text string `json:"text,omitempty"`
	// Estimation of the carbon dioxide emission for the given Route. The value depends on the VehicleType
	// request parameter which specifies the average fuel consumption per 100 km and the type of combustion
	// engine (diesel or gasoline). Unit is kilograms with precision to three decimal places.
	CO2Emission json.RawMessage `json:"co2Emission,omitempty"`
}

// RouteLeg is the portion of a route between one waypoint and the next.
//
// A RouteLeg contains information about a route leg, such as the time required to traverse it, its shape, start point
// and endpoint, as well as information about any sublegs contained in the leg due to the presence of passthrough
// waypoints.
type RouteLeg struct {
	// LengthMeters is the length of the leg.
	LengthMeters float64 `json:"length,omitempty"`

	// Links is a list of all links which are included in this portion of the route.
	Links []RouteLink `json:"link,omitempty"`

	// BaseTime is the estimated time spent on this leg, without considering traffic conditions.
	// The service may also account for additional time penalties, therefore this may be greater than the leg length
	// divided by the base speed.
	BaseTime Duration `json:"baseTime,omitempty"`
}

// RouteLink is a path segment in the routing network, such as a road.
type RouteLink struct {
	// ID is a permanent ID which references a network link.
	// When presented with a minus sign as the first character, this ID indicates that the link should be traversed in
	// the opposite direction of its default coding (for example, walking SW on a link that is coded as one-way
	// traveling NE).
	ID string `json:"linkId,omitempty"`

	// Shapes is a polyline representation of the link.
	Shape []LatLng `json:"shape,omitempty"`

	// LengthMeters is the length of the link.
	LengthMeters float64 `json:"length,omitempty"`

	// RemainingDistanceMeters from the start of this element to the destination of the route.
	RemainingDistanceMeters float64 `json:"remainDistance,omitempty"`

	// RemainingTimeMeters needed from the start of this element to the destination of the route.
	// Considers any available traffic information, if enabled and the authorized for the user.
	RemainingTime Duration `json:"remainTime,omitempty"`
}

type RouteLinkFlag string

const (
	// Link can only be traversed by using a boat ferry
	RouteLinkFlagBoatFerry = "ferry"
	// Link is part of a dirt road
	RouteLinkFlagDirtRoad = "dirtRoad"
	// Link can only be traversed by using high-occupancy vehicle (HOV) lanes
	RouteLinkFlagHOVLane = "HOVLane"
	// Link is part of a motorway
	RouteLinkFlagMotorway = "motorway"
	// Link is part of a road that you can enter but you have to exit the same way
	RouteLinkFlagNoThroughRoad = "noThroughRoad"
	// Link is part of a park
	RouteLinkFlagPark = "park"
	// Link is part of a private road
	RouteLinkFlagPrivateRoad = "privateRoad"
	// Link can only be traversed by using a rail ferry
	RouteLinkFlagRailFerry = "railFerry"
	// Link is part of a toll road
	RouteLinkFlagTollRoad = "tollRoad"
	// Link passes through a tunnel
	RouteLinkFlagTunnel = "tunnel"
	// Link is part of a built-up area
	RouteLinkFlagBuiltUpArea = "builtUpArea"
)
