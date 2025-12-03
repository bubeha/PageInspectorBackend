package validator

import (
	"github.com/bubeha/PageInspectorBackend/pkg/log"
	"github.com/go-playground/validator/v10"
)

var (
	Instance = validator.New()
)

func Init() {
	if err := Instance.RegisterValidation("domain", validateDomain); err != nil {
		log.Errorf("validator init error: %v", err)
	}
}

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
