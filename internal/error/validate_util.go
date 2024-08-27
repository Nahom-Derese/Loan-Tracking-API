package error_handler

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// TranslateError translates validator errors into custom messages
func TranslateError(err error) map[string]string {
	errs := err.(validator.ValidationErrors)
	errorMessages := make(map[string]string)

	for _, e := range errs {
		fieldName := e.Field()
		tag := e.Tag()
		var message string

		switch tag {
		case "required":
			message = fmt.Sprintf("%s is required", fieldName)
		case "email":
			message = fmt.Sprintf("%s must be a valid email address", fieldName)
		case "min":
			message = fmt.Sprintf("%s must be at least %s characters long", fieldName, e.Param())
		case "max":
			message = fmt.Sprintf("%s cannot be longer than %s characters", fieldName, e.Param())
		case "e164":
			message = fmt.Sprintf("%s must be a valid phone number in E.164 format", fieldName)
		default:
			message = fmt.Sprintf("%s is not valid", fieldName)
		}

		errorMessages[fieldName] = message
	}

	return errorMessages
}
