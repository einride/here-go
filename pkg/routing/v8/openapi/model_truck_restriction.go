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
// TruckRestriction Contains:  * merged information about all truck restrictions applicable independent of truck properties or time.   Examples are weight limit, height limit etc. * truck restrictions which depend on truck parameters or time.   Examples are weight limit depending on number of axles, access restrictions depending on time etc. 
type TruckRestriction struct {
	// Hazardous goods restrictions applied during the trip.
	ForbiddenHazardousGoods []string `json:"forbiddenHazardousGoods,omitempty"`
	// Contains max permitted gross weight. 
	MaxGrossWeight int32 `json:"maxGrossWeight,omitempty"`
	// Contains max permitted weight per axle. 
	MaxWeightPerAxle int32 `json:"maxWeightPerAxle,omitempty"`
	// Contains max permitted height. 
	MaxHeight int32 `json:"maxHeight,omitempty"`
	// Contains max permitted width. 
	MaxWidth int32 `json:"maxWidth,omitempty"`
	// Contains max permitted length. 
	MaxLength int32 `json:"maxLength,omitempty"`
	// Contains max permitted axle count. 
	MaxAxleCount int32 `json:"maxAxleCount,omitempty"`
	TunnelCategory TunnelCategory `json:"tunnelCategory,omitempty"`
	// Indicates that restriction depends on time. 
	TimeDependent bool `json:"timeDependent,omitempty"`
	AppliesIf TruckProfileFilter `json:"appliesIf,omitempty"`
}
