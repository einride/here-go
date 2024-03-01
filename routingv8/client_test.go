package routingv8_test

import (
	"testing"

	"go.einride.tech/here/routingv8"
	"gotest.tools/v3/assert"
)

func TestRouteError(t *testing.T) {
	for _, tt := range []*struct {
		name                string
		expectedErrorString string
		responseError       routingv8.ResponseError
	}{
		{
			name: "normally serialized",
			responseError: routingv8.ResponseError{
				Response: &routingv8.HereErrorResponse{
					Title:  "Unknown",
					Status: 500,
					Code:   "ErrorCode",
					Cause:  "Random Cause",
					Action: "Do it right",
				},
			},
			expectedErrorString: "Title: Unknown, Status: 500, Code: ErrorCode, Cause: Random Cause, Action: Do it right",
		},
		{
			name: "response not present",
			responseError: routingv8.ResponseError{
				HTTPBody: "{}",
			},
			expectedErrorString: "Response: {} StatusCode: 0",
		},
		{
			name: "response malformed",
			responseError: routingv8.ResponseError{
				Response: &routingv8.HereErrorResponse{},
				HTTPBody: "{}",
			},
			expectedErrorString: "Response: {} StatusCode: 0",
		},
		{
			name: "response with 0 status code",
			responseError: routingv8.ResponseError{
				Response: &routingv8.HereErrorResponse{
					Status: 0,
				},
				HTTPBody: "{}",
			},
			expectedErrorString: "Response: {} StatusCode: 0",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.responseError.Error(), tt.expectedErrorString)
		})
	}
}
