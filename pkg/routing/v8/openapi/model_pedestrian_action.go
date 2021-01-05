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
// PedestrianAction Action attached to a pedestrian section.
type PedestrianAction struct {
	// The type of the action.  **NOTE:** The list of possible actions may be extended in the future. The client application should handle such a case gracefully. 
	Action string `json:"action"`
	// Estimated duration of this action. Actions last until the next action, or the end of the route in case of the last one.
	Duration int32 `json:"duration"`
	// Description of the action (e.g. Turn left onto Minna St.).
	Instruction string `json:"instruction,omitempty"`
	// Offset of a coordinate in the section's polyline.
	Offset int32 `json:"offset,omitempty"`
	// Attributes of the current road
	CurrentRoad RoadInfo `json:"currentRoad,omitempty"`
	// Attributes of the next road
	NextRoad RoadInfo `json:"nextRoad,omitempty"`
	Direction TurnActionDirection `json:"direction,omitempty"`
	Severity TurnActionSeverity `json:"severity,omitempty"`
	// Which exit to take next.
	Exit int32 `json:"exit,omitempty"`
}
