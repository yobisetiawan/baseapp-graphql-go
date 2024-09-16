package utils

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}

type ValidationErrorResponse struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}

func NewValidationErrorResponse(err error) *ValidationErrorResponse {
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		errorsMap := make(map[string]string)

		for _, fe := range ve {

			var sTag = fe.Tag()
			if sTag == "required" {
				sTag = "This field is required"
			}
			if sTag == "min" {
				sTag = "This field is too short"
			}
			if sTag == "url" {
				sTag = "Invalid URL"
			}
			if sTag == "email" {
				sTag = "Invalid Email"
			}

			errorsMap[StringToSnakeCase(fe.Field())] = sTag
		}
		return &ValidationErrorResponse{
			Message: "Validation failed",
			Errors:  errorsMap,
		}
	}

	// Handle other types of errors if necessary
	return &ValidationErrorResponse{
		Message: "Invalid request",
		Errors:  map[string]string{"error": err.Error()},
	}
}
