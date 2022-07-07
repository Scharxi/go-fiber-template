package error

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type (
	ResponseError struct {
		Status     int       `json:"status"`
		Message    string    `json:"error"`
		OccurredOn time.Time `json:"occurredOn"`
	}

	ValidationErrorResponse struct {
		ResponseError
		FailedField string `json:"failedField"`
		Tag         string `json:"tag"`
		Value       string `json:"value"`
	}
)

func NewErrorResponse(status int, message string) *ResponseError {
	return &ResponseError{
		Status:     status,
		Message:    message,
		OccurredOn: time.Now(),
	}
}

func NewValidationErrorResponse(err error) []*ValidationErrorResponse {
	var errs []*ValidationErrorResponse
	for _, err := range err.(validator.ValidationErrors) {
		element := ValidationErrorResponse{
			ResponseError: ResponseError{
				Status:     400,
				Message:    "Missing required fields",
				OccurredOn: time.Now(),
			},
			FailedField: err.StructNamespace(),
			Tag:         err.Tag(),
			Value:       err.Param(),
		}
		errs = append(errs, &element)
	}
	return errs
}

var (
	InvalidCredentialsError = ResponseError{
		Status:     401,
		Message:    "Invalid credentials",
		OccurredOn: time.Now(),
	}
	NotFoundError = ResponseError{
		Status:     404,
		Message:    "Not found",
		OccurredOn: time.Now(),
	}
	InternalServerError = ResponseError{
		Status:     500,
		Message:    "Internal server error",
		OccurredOn: time.Now(),
	}
	InvalidRequestBodyError = ResponseError{
		Status:     400,
		Message:    "Invalid request body",
		OccurredOn: time.Now(),
	}
)
