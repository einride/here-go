// Code generated by OpenAPI Generator. DO NOT EDIT.
/*
 * Routing API v8
 *
 * A location service providing customizable route calculations for a variety of vehicle types as well as pedestrian modes.
 *
 * API version: 8.3.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package routingv8
// PedestrianPlace Place used in pedestrian routing
type PedestrianPlace struct {
	// Location name
	Name string `json:"name,omitempty"`
	// If present, this place corresponds to the waypoint in the request with the same index.
	Waypoint int32 `json:"waypoint,omitempty"`
	// Place type. Each place type can have extra attributes.  **NOTE:** Please note that the list of possible place types could be extended in the future. The client application is expected to handle such a case gracefully. 
	Type string `json:"type"`
	// The position of this location  This position was used in route calculation. It may be different to the original position provided in the request. 
	Location Location `json:"location"`
	// If present, the original position of this location provided in the request.
	OriginalLocation Location `json:"originalLocation,omitempty"`
	// Identifier of this charging station
	Id string `json:"id,omitempty"`
	// Platform name or number for the departure.
	Platform string `json:"platform,omitempty"`
	// Attributes of a parking lot.  * `parkAndRide` - this parking lot is of type \"Park and Ride\",   such as it is a parking specifically designed to allow transition between car and transit. 
	Attributes []string `json:"attributes,omitempty"`
	// List of possible parking rates for this facility. Different rates can apply depending on the day, time of the day or parking duration. 
	Rates []TimeRestrictedPrice `json:"rates,omitempty"`
}
