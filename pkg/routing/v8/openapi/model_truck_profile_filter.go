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
// TruckProfileFilter Specifies which truck profiles a restriction applies to. 
type TruckProfileFilter struct {
	AxleCount TruckAxleCountRange `json:"axleCount,omitempty"`
}
