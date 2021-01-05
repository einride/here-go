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
// AccessRestrictionNoticeDetail Contains details about violated access restriction. 
type AccessRestrictionNoticeDetail struct {
	// Detail title
	Title string `json:"title,omitempty"`
	// Cause of the notice
	Cause string `json:"cause,omitempty"`
	// Detail type. Each type of detail might contain extra attributes.  **NOTE:** The list of possible detail types may be extended in the future. The client application is expected to handle such a case gracefully. 
	Type string `json:"type"`
}
