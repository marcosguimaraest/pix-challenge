package objects

import "encoding/json"

type ErrorResponse struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type ArrayErrorResponse struct {
	Errors []ErrorResponse `json:"errors"`
}

func ByteToErrorResponse(b []byte) (ErrorResponse, error) {
	var errorResponse ErrorResponse
	err := json.Unmarshal(b, &errorResponse)
	return errorResponse, err
}

func ByteToErrorResponseArray(b []byte) (ArrayErrorResponse, error) {
	var arrayErrorResponse ArrayErrorResponse
	err := json.Unmarshal(b, &arrayErrorResponse)
	return arrayErrorResponse, err
}
