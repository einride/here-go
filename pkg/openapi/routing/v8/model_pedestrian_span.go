/*
 * Routing API v8
 *
 * A location service providing customizable route calculations for a variety of vehicle types as well as pedestrian modes.
 *
 * API version: 8.3.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package routingv8
// PedestrianSpan Span attached to a `Section` describing pedestrian content. 
type PedestrianSpan struct {
	// Offset of a coordinate in the section's polyline. 
	Offset int32 `json:"offset,omitempty"`
	// Distance in meters
	Length int32 `json:"length,omitempty"`
	// Duration in seconds
	Duration int32 `json:"duration,omitempty"`
	// `StreetAttributes` is applied to a span of a route section and describes attribute flags of a street. * `rightDrivingSide`: Do vehicles have to drive on the right-hand side of the road or the left-hand side. * `dirtRoad`: This part of the route has an un-paved surface. * `tunnel`: This part of the route is a tunnel. * `bridge`: This part of the route is a bridge. * `ramp`: This part of the route is a ramp (usually connecting to/from/between highways). * `motorway`: This part of the route is a controlled access road (usually highways). * `roundabout`: This part of the route is a roundabout. * `underConstruction`: This part of the route is under construction. * `dividedRoad`: This part of the route uses a road with a physical or legal divider in the middle. * `privateRoad`: This part of the route uses a privately owned road.  As it is possible that new street attributes are supported in the future, unknown street attributes should be ignored. 
	StreetAttributes []string `json:"streetAttributes,omitempty"`
	// Accessibility and walk-related attribute flags.  * `stairs`: This part of the route is a stair case. * `park`: This part of the route is in a park. * `indoor`: This part of the route is inside a venue. * `open`:  Describes whether this part of the route can be traversed. * `noThrough`:  A part of the route you can only enter if your destination is located there. * `tollRoad`: Access to this part of the route is restricted with a fee (or toll).  As it is possible that new attributes are supported in the future, unknown attributes should be ignored. 
	WalkAttributes []string `json:"walkAttributes,omitempty"`
	// Car specific `AccessAttributes`.  `AccessAttributes` is applied to a span of a route section and describes access flags of a street. * `open`:  Describes if you are allowed to traverse this stretch of the route. * `noThrough`:  A part of the route you can only enter if your destination is located there. * `tollRoad`: Access to this part of the route is restricted with a fee (or toll).  As it is possible that new access attributes are supported in the future, unknown access attributes should be ignored. 
	CarAttributes []string `json:"carAttributes,omitempty"`
	// Designated name for the span (e.g. a street name or a transport name)
	Names []LocalizedString `json:"names,omitempty"`
	// Designated route name or number of the span (e.g. 'M25')
	RouteNumbers []LocalizedString `json:"routeNumbers,omitempty"`
	// ISO-3166-1 alpha-3 code
	CountryCode string `json:"countryCode,omitempty"`
	// Functional class is used to classify roads depending on the speed, importance and connectivity of the road.  * `1`: Roads allow for high volume, maximum speed traffic movement between and through major   metropolitan areas. * `2`: Roads are used to channel traffic to functional class 1 roads for travel between and   through cities in the shortest amount of time. * `3`: Roads that intersect functional class 2 roads and provide a high volume of traffic   movement at a lower level of mobility than functional class 2 roads. * `4`: Roads that provide for a high volume of traffic movement at moderate speeds between   neighbourhoods. * `5`: Roads with volume and traffic movement below the level of any other functional class. 
	FunctionalClass int32 `json:"functionalClass,omitempty"`
	// Speed in meters per second
	SpeedLimit float32 `json:"speedLimit,omitempty"`
	DynamicSpeedInfo DynamicSpeedInfo `json:"dynamicSpeedInfo,omitempty"`
}
