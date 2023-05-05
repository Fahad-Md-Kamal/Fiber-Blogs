package utils

import (
	"fmt"
	"log"

	"github.com/go-playground/validator"
)

type ErrorResponse struct {
	FailedField string `json:"field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

var validate = validator.New()

func ValidateStruct(inputStruct interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(inputStruct)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	var errorMsg string
	if len(errors) > 0 {
		for _, err := range errors {
			errorMsg += fmt.Sprintf("Field: %s Tag: %s Value: %s", err.FailedField, err.Tag, err.Value) + "\n"
		}
		log.Printf("Invalid User data. Errors: %s", errorMsg)
	}
	return errors
}
