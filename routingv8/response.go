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
