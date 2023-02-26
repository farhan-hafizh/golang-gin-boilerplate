package helper

import "github.com/go-playground/validator/v10"

func FormatValidationError(err error) []string {
	var errors []string

	// loop through error that had changed to ValidationErrors type
	// then append the error to errors
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
