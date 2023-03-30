package routingv8

import (
	"encoding/json"
)

type ErrorCodes []ErrorCode

func (e *ErrorCodes) UnmarshalJSON(b []byte) error {
	errorCodes := make([]int, 0)
	if err := json.Unmarshal(b, &errorCodes); err != nil {
		return err
	}
	codes := make([]ErrorCode, 0, len(errorCodes))
	for _, errorCode := range errorCodes {
		codes = append(codes, ErrorCode(errorCode))
	}
	*e = codes
	return nil
}

type ErrorCode int

// See https://developer.here.com/documentation/matrix-routing-api/8.6.0/api-reference-swagger.html
// for detailed explanations of each error.
const (
	ErrorCodeSuccess            = 0
	ErrorCodeDisconnected       = 1
	ErrorCodeMatchingFailed     = 2
	ErrorCodeParameterViolation = 3
	ErrorCodeUnknown            = 99
)

// MatrixResponse contains the calculated route matrix.
type MatrixResponse struct {
	// NumOrigins used to calculate matrix
	NumOrigins int `json:"numOrigins"`
	// NumDestinations used to calculate matrix
	NumDestinations int `json:"numDestinations"`
	// TravelTimes calculated using origins and destinations. Nil if not requested in MatrixAttributes.
	TravelTimes []int32 `json:"travelTimes"`
	// Distances calculated using origins and destinations. Nil if not requested in MatrixAttributes.
	Distances []int32 `json:"distances"`
	// ErrorCodes contains potential route errors. Nil if no errors occurred.
	ErrorCodes ErrorCodes `json:"errorCodes"`
}

// CalculateMatrixResponse is used to provide results of a matrix calculation.
type CalculateMatrixResponse struct {
	// MatrixID is unique identifier of the matrix
	MatrixID string `json:"matrixId"`
	// Matrix contains the calculated matrix data.
	Matrix MatrixResponse `json:"matrix"`
	// RegionDefinition to be used to calculate matrix.
	RegionDefinition RegionDefinition `json:"regionDefinition"`
}

// RoutesResponse contains the possible routes.
type RoutesResponse struct {
	// Routes in the possible routes between the origin and target.
	Routes []Route `json:"routes"`
	// Contains a list of issues related to this route calculation.
	Notices []RouteResponseNotice `json:"notices"`
}

// Route contains all the sections of a route.
type Route struct {
	// ID of the route
	ID string `json:"id"`
	// Sections in the route
	Sections []Section `json:"sections"`
	// Contains a list of issues related to this route calculation.
	Notices []Notice `json:"notices"`
}

type RouteResponseNotice struct {
	// Human-readable notice description.
	Title string `json:"title"`
	// Machine-readable notice code.
	// See https://developer.here.com/documentation/routing-api/api-reference-swagger.html
	// for possible values.
	Code string `json:"code"`
	// Describes the impact a notice has on the resource to which the notice is attached.
	Severity NoticeSeverity `json:"severity"`
}

type Notice struct {
	// Human-readable notice description.
	Title string `json:"title"`
	// Machine-readable notice code.
	// See https://developer.here.com/documentation/routing-api/api-reference-swagger.html
	// for possible values.
	Code string `json:"code"`
	// Describes the impact a notice has on the resource to which the notice is attached.
	Severity NoticeSeverity `json:"severity"`
}

type NoticeSeverity string

const (
	// CriticalNoticeSeverity is used to indicate that the notice must not be ignored,
	// even if the type of notice is not known to the user.
	// Any associated resource (e.g., route section) must not be used without further evaluation.
	CriticalNoticeSeverity NoticeSeverity = "critical"
	// InfoNoticeSeverity is used to indicate that the notice is for informative purposes,
	// but does not affect usability of the route.
	InfoNoticeSeverity NoticeSeverity = "info"
)

// Section with the information of the departure, arrival location and summary.
type Section struct {
	// ID of the section
	ID string `json:"id"`
	// The type used in the section
	Type string `json:"type"`
	// Departure is the location of the departure
	Departure Place `json:"departure"`
	// Arrival is the location of the arrival
	Arrival Place `json:"arrival"`
	// Summary contain info on the duration and length of section
	Summary Summary `json:"summary"`
}

// Place with lat and long info on where the place is.
type Place struct {
	// Type is the struct
	Type string `json:"type"`
	// Location in lat and long
	Location GeoWaypoint `json:"location"`
	// OriginalLocation in lat and long
	OriginalLocation GeoWaypoint `json:"originalLocation"`
}

// Summary contains the duration and length info.
type Summary struct {
	// Duration is the total duration of the action, section etc
	Duration int32 `json:"duration"`
	// Length is the total length
	Length int32 `json:"length"`
	// BaseDuration is the duration without dynamic traffic information
	BaseDuration int32 `json:"baseDuration"`
}

// HereErrorResponse is returned when an error is returned from the Here Maps API.
type HereErrorResponse struct {
	// Title of the error
	Title string `json:"title"`
	// Http status code
	Status int `json:"status"`
	// Here Maps API error code
	Code string `json:"code"`
	// Cause of the error
	Cause string `json:"cause"`
	// Action Suggested to fix error
	Action string `json:"action"`
}
