package routingv7

import "encoding/json"

type RouteResponseMetaInfo struct {
	// Mirrored RequestId value from the request structure. Used to trace requests.
	RequestID string `json:"requestId,omitempty"`
	// Time at which the search was performed.
	Timestamp string `json:"timestamp,omitempty"`
	// Gives the version of the underlying map, upon which the route calculations are based.
	MapVersion string `json:"mapVersion,omitempty"`
	// Gives a list of map versions, upon which the route calculations are based.
	AvailableMapVersion []string `json:"availableMapVersion,omitempty"`
	// Gives the version of the module that performed the route calculations.
	ModuleVersion string `json:"moduleVersion,omitempty"`
	// Required. Gives the version of the schema definition to enable
	InterfaceVersion string `json:"interfaceVersion,omitempty"`
}

type Route struct {
	// Overall route distance and time summary.
	Summary RouteSummary `json:"summary,omitempty"`
}

type RouteSummary struct {
	// Indicates total travel distance for the route, in meters.
	Distance float64 `json:"distance,omitempty"`
	// Contains the travel time estimate in seconds for this element, considering traffic
	// and transport mode. Based on the TrafficSpeed. The service may also account for additional
	// time penalties, so this may be greater than the element length divided by the TrafficSpeed.
	TrafficTime float64 `json:"traffictime,omitempty"`
	// Contains the travel time estimate in seconds for this element, considering transport mode but
	// not traffic conditions. Based on the BaseSpeed. The service may also account for additional
	// time penalties, therefore this may be greater than the element length divided by the BaseSpeed.
	BaseTime float64 `json:"baseTime,omitempty"`
	// Special link characteristics (like ferry or motorway usage) which are covered by the route.
	Flags []RouteLinkFlag `json:"flags,omitempty"`
	// Total travel time in seconds optionally considering traffic depending on the request parameters.
	TravelTime float64 `json:"travelTime,omitempty"`
	// Textual description of route summary. Supported languages are US and GB English, French, German,
	// Italian, Polish, Portuguese, Spanish and Traditional Chinese. Element is not provided in case of
	// other languages.
	Text string `json:"text,omitempty"`
	// Estimation of the carbon dioxide emission for the given Route. The value depends on the VehicleType
	// request parameter which specifies the average fuel consumption per 100 km and the type of combustion
	// engine (diesel or gasoline). Unit is kilograms with precision to three decimal places.
	CO2Emission json.RawMessage `json:"co2Emission,omitempty"`
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
