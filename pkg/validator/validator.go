package validator

import (
	"github.com/go-playground/validator/v10"
)

var (
	Instance = validator.New()
)

func Validate(i interface{}) error {
	if err := Instance.Struct(i); err != nil {
		validationErrors := make(map[string][]string)

		for _, e := range err.(validator.ValidationErrors) {
			field := e.Field()

			validationErrors[field] = append(validationErrors[field], formatValidationError(e))
		}

		return ValidationError{Errors: validationErrors}
	}

	return nil
}
