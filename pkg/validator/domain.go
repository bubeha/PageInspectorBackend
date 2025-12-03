package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var domainRegex = regexp.MustCompile(`^(?:[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`)

func validateDomain(fl validator.FieldLevel) bool {
	domain := fl.Field().String()

	return domainRegex.MatchString(domain)
}
