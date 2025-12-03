package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Errors map[string][]string `json:"errors"`
}

func (e ValidationError) Error() string {
	var messages []string

	for field, msg := range e.Errors {
		messages = append(messages, fmt.Sprintf("%s: %s", field, strings.Join(msg, ";")))
	}

	return strings.Join(messages, "")
}

func formatValidationError(e validator.FieldError) string {
	tag := e.Tag()
	fieldName := e.Field()
	param := e.Param()

	switch tag {
	case "required":
		return fmt.Sprintf("%s is required", fieldName)
	case "required_if", "required_unless", "required_with", "required_without", "required_with_all", "required_without_all":
		return fmt.Sprintf("%s is required under these conditions", fieldName)
	case "email":
		return fmt.Sprintf("%s must be a valid email address", fieldName)

	case "len":
		return fmt.Sprintf("%s must be exactly %s characters long", fieldName, param)

	case "min":
		if isNumericField(e.Kind()) {
			return fmt.Sprintf("%s must be at least %s", fieldName, param)
		}

		return fmt.Sprintf("%s must be at least %s characters long", fieldName, param)

	case "max":
		if isNumericField(e.Kind()) {
			return fmt.Sprintf("%s must be at most %s", fieldName, param)
		}
		return fmt.Sprintf("%s must be at most %s characters long", fieldName, param)

	case "eq":
		return fmt.Sprintf("%s must be equal to %s", fieldName, param)

	case "ne":
		return fmt.Sprintf("%s must not be equal to %s", fieldName, param)

	case "gt":
		return fmt.Sprintf("%s must be greater than %s", fieldName, param)

	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %s", fieldName, param)

	case "lt":
		return fmt.Sprintf("%s must be less than %s", fieldName, param)

	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %s", fieldName, param)

	case "oneof":
		options := strings.ReplaceAll(param, " ", ", ")
		return fmt.Sprintf("%s must be one of: %s", fieldName, options)

	case "unique":
		return fmt.Sprintf("%s must contain unique values", fieldName)

	case "alphanum":
		return fmt.Sprintf("%s must contain only letters and numbers", fieldName)

	case "alpha":
		return fmt.Sprintf("%s must contain only letters", fieldName)

	case "numeric":
		return fmt.Sprintf("%s must be a valid number", fieldName)

	case "hexadecimal":
		return fmt.Sprintf("%s must be a valid hexadecimal number", fieldName)

	case "hexcolor":
		return fmt.Sprintf("%s must be a valid HEX color code", fieldName)

	case "rgb", "rgba":
		return fmt.Sprintf("%s must be a valid RGB/RGBA color", fieldName)

	case "hsl", "hsla":
		return fmt.Sprintf("%s must be a valid HSL/HSLA color", fieldName)

	case "e164":
		return fmt.Sprintf("%s must be a valid E.164 phone number", fieldName)

	case "uuid", "uuid3", "uuid4", "uuid5":
		return fmt.Sprintf("%s must be a valid UUID", fieldName)

	case "url":
		return fmt.Sprintf("%s must be a valid URL", fieldName)

	case "uri":
		return fmt.Sprintf("%s must be a valid URI", fieldName)

	case "ip", "ipv4", "ipv6":
		return fmt.Sprintf("%s must be a valid IP address", fieldName)

	case "mac":
		return fmt.Sprintf("%s must be a valid MAC address", fieldName)

	case "latitude":
		return fmt.Sprintf("%s must be a valid latitude coordinate", fieldName)

	case "longitude":
		return fmt.Sprintf("%s must be a valid longitude coordinate", fieldName)

	case "datetime":
		if param != "" {
			return fmt.Sprintf("%s must be in %s format", fieldName, param)
		}
		return fmt.Sprintf("%s must be a valid date and time", fieldName)

	case "date":
		return fmt.Sprintf("%s must be a valid date", fieldName)

	case "boolean":
		return fmt.Sprintf("%s must be a boolean value", fieldName)

	case "json":
		return fmt.Sprintf("%s must be a valid JSON string", fieldName)

	case "file":
		return fmt.Sprintf("%s must be a valid file", fieldName)

	case "image":
		return fmt.Sprintf("%s must be a valid image", fieldName)

	case "lowercase":
		return fmt.Sprintf("%s must be in lowercase", fieldName)

	case "uppercase":
		return fmt.Sprintf("%s must be in uppercase", fieldName)

	case "password":
		return fmt.Sprintf("%s must contain at least one uppercase letter, one lowercase letter, one number, and one special character", fieldName)

	case "strong_password":
		return fmt.Sprintf("%s must contain at least one uppercase letter, one lowercase letter, one number, and one special character", fieldName)

	case "confirmed":
		return fmt.Sprintf("%s confirmation does not match", fieldName)

	case "different":
		return fmt.Sprintf("%s must be different from %s", fieldName, param)

	case "same":
		return fmt.Sprintf("%s must be the same as %s", fieldName, param)

	case "startsnotwith", "endsnotwith":
		return fmt.Sprintf("%s must not start/end with %s", fieldName, param)

	case "starts_with", "ends_with":
		return fmt.Sprintf("%s must start/end with %s", fieldName, param)

	case "domain":
		return fmt.Sprintf("%s must be a valid domain name", fieldName)

	default:
		return fmt.Sprintf("%s has an invalid value", fieldName)
	}
}

func isNumericField(kind reflect.Kind) bool {
	numericKinds := []reflect.Kind{
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
	}

	for _, k := range numericKinds {
		if kind == k {
			return true
		}
	}

	return false
}
