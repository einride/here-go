package routingv8

import (
	"fmt"
	"strconv"
)

type ErrorCode int

// See https://developer.here.com/documentation/matrix-routing-api/8.6.0/api-reference-swagger.html
// for detailed explanations of each error.
const (
	ErrorCodeUnspecified ErrorCode = iota
	ErrorCodeSuccess
	ErrorCodeDisconnected
	ErrorCodeMatchingFailed
	ErrorCodeParameterViolation
	ErrorCodeUnknown
)

func (e *ErrorCode) UnmarshalString(value string) error {
	switch value {
	case "0":
		*e = ErrorCodeSuccess
	case "1":
		*e = ErrorCodeDisconnected
	case "2":
		*e = ErrorCodeMatchingFailed
	case "3":
		*e = ErrorCodeParameterViolation
	case "99":
		*e = ErrorCodeUnknown
	default:
		return fmt.Errorf("invalid error code")
	}
	return nil
}

func (e *ErrorCode) UnmarshalJSON(b []byte) error {
	value, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	return e.UnmarshalString(value)
}

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
	ErrorCodes []ErrorCode `json:"errorCodes"`
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
