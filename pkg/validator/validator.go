package validator

import "github.com/go-playground/validator/v10"

var (
	Instance = validator.New()
)

func Validate(i interface{}) error {
	return Instance.Struct(i)
}
