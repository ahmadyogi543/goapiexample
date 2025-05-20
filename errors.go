package main

import (
	"fmt"
	"net/http"
)

type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

type APIErrorResponse struct {
	Errors APIError `json:"errors"`
}

func (er APIErrorResponse) Error() string {
	return fmt.Sprintf("api error: %s", er.Errors.Message)
}

func NewAPIResponse(status int, message string, details any) APIErrorResponse {
	return APIErrorResponse{
		Errors: APIError{
			Status:  status,
			Message: message,
			Details: details,
		},
	}
}

func ErrBadRequestResponse(message string, details any) APIErrorResponse {
	return NewAPIResponse(http.StatusBadRequest, message, details)
}

func ErrNotFoundResponse(message string) APIErrorResponse {
	return NewAPIResponse(http.StatusNotFound, message, nil)
}
