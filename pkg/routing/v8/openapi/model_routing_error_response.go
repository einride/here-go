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
// RoutingErrorResponse Response in case of error
type RoutingErrorResponse struct {
	// Human readable error description
	Title string `json:"title"`
	// HTTP status code
	Status int32 `json:"status"`
	// Machine readable service error code.  All error codes of this service start with \"`E605`\". The last three digits describe a specific error. Provide this error code when contacting support.  **NOTE:** Please note that the list of possible error codes could be extended in the future. The client application is expected to handle such a case gracefully.  | Code      | Reason  | | --------- | ------- | | `E60500X` | Malformed query. Typically due to invalid values such as `transportMode=spaceShuttle` or missing required fields. Check the error message for details. | | `E605010` | Invalid combination of truck and transport mode. Check `truck` for valid truck transport modes. | | `E605011` | Invalid combination of avoid feature `difficultTurns` and transport mode. Check `avoid` for details. | | `E605012` | Invalid combination of transport mode and routing mode. Check `routingMode` for a list of supported combinations. | | `E605013` | Invalid return options. Check `return` for valid combinations of values. | | `E605014` | Invalid language code. Check `lang` for details on how valid language codes look. | | `E605015` | Too many alternatives. Check `alternatives` for the maximum number of alternatives allowed. | | `E605016` | Invalid exclude countries. Check `exclude` for details. | | `E60503X` | Invalid EV options. Check `ev` for details on how valid EV options look. | | `E605040` | Invalid combination of EV and transport mode. Check `ev` for details. | | `E605041` | Invalid combination of EV and routing mode. Check `ev` for details. | | `E605042` | Invalid combination of EV and alternatives. Check `ev` for details. | | `E605043` | Invalid combination of EV and avoid options. Check `ev` for details. | | `E605044` | Invalid combination of EV and exclude options. Check `ev` for details. | | `E605101` | Credentials not allowed for calculating routes in Japan. | | `E6055XX` | Internal server error. | 
	Code string `json:"code"`
	// Human readable explanation for the error
	Cause string `json:"cause"`
	// Human readable action that can be taken to correct the error
	Action string `json:"action"`
	// Auto generated id that univocally identify this request
	CorrelationId string `json:"correlationId"`
}
