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
// SinglePrice struct for SinglePrice
type SinglePrice struct {
	// Type of price represented by this object. The API customer is responsible for correctly visualizing the pricing model. As it is possible to extend the supported price types in the future, the price information should be hidden when an unknown type is encountered.  Available price types are:    * `value` - A single value.   * `range` - A range value that includes a minimum and maximum price. 
	Type string `json:"type"`
	// Attribute value is `true` if the fare price is estimated, `false` if it is an exact value.
	Estimated bool `json:"estimated,omitempty"`
	// Local currency of the price compliant to ISO 4217
	Currency string `json:"currency"`
	// When set, the price is paid for a specific duration.  Examples:   * `\"unit\": 3600` - price for one hour   * `\"unit\": 28800` - price for 8 hours   * `\"unit\": 86400` - price for one day 
	Unit int32 `json:"unit,omitempty"`
	// The price value
	Value float32 `json:"value"`
}
