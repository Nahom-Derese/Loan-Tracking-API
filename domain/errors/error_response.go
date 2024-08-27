package custom_error

import (
	"errors"
	"net/http"
)

type ErrorResponse struct {
	Error      ErrorMessage              `json:"error,omitempty"`
	Validation []ValidationErrorResponse `json:"validation,omitempty"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

func ErrMessage(message error) ErrorResponse {
	return ErrorResponse{
		Error: ErrorMessage{
			Message: message.Error(),
		},
	}
}

func ErrValidation(validation []ValidationErrorResponse) ErrorResponse {
	return ErrorResponse{
		Validation: validation,
	}
}

type ValidationErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// General Errors
var (
	ErrInvalidToken      = errors.New("invalid token")
	ErrUnauthorized      = errors.New("unauthorized")
	ErrUserNotFound      = errors.New("user not found")
	ErrUserNotVerified   = errors.New("user is not verified")
	ErrAlreadyVerified   = errors.New("user is already verified")
	ErrUserAlreadyExists = errors.New("user already exists with the given email")

	ErrUserNotActive                 = errors.New("user is not active")
	ErrInvalidPasswordLength         = errors.New("password must be at least 8 characters long")
	ErrErrorBindingRequest           = errors.New("error binding request")
	ErrErrorEncryptingPassword       = errors.New("error encrypting password")
	ErrErrorCreatingUser             = errors.New("error creating user")
	ErrErrorSendingVerificationEmail = errors.New("error sending verification email")
	ErrInvalidEmailFormat            = errors.New("invalid email format")
	EreInvalidRequestBody            = errors.New("request body cannot be empty")

	ErrRateLimitExceeded         = errors.New("rate limit exceeded")
	ErrErrorSendingReminderEmail = errors.New("error sending reminder email")

	ErrInvalidID           = errors.New("invalid id")
	ErrFilteringUsers      = errors.New("error filtering users")
	ErrTokenNotFound       = errors.New("token not found")
	ErrErrorUpdatingUser   = errors.New("error updating user")
	ErrCredentialsNotValid = errors.New("credentials not valid")
	ErrErrorCreatingLoan   = errors.New("error creating loan")
	ErrLoanNotFound        = errors.New("loan not found")
	ErrInvalidLoanAmount   = errors.New("invalid amount of loan")
)

var errorStatusMap = map[error]int{
	ErrInvalidToken:                  http.StatusUnauthorized,
	ErrUnauthorized:                  http.StatusUnauthorized,
	ErrUserNotFound:                  http.StatusNotFound,
	ErrUserNotVerified:               http.StatusUnauthorized,
	ErrAlreadyVerified:               http.StatusBadRequest,
	ErrUserAlreadyExists:             http.StatusBadRequest,
	ErrUserNotActive:                 http.StatusForbidden,
	ErrInvalidPasswordLength:         http.StatusBadRequest,
	ErrErrorBindingRequest:           http.StatusBadRequest,
	ErrErrorEncryptingPassword:       http.StatusInternalServerError,
	ErrErrorCreatingUser:             http.StatusInternalServerError,
	ErrErrorSendingVerificationEmail: http.StatusInternalServerError,
	ErrInvalidEmailFormat:            http.StatusBadRequest,
	EreInvalidRequestBody:            http.StatusBadRequest,
	ErrRateLimitExceeded:             http.StatusTooManyRequests,
	ErrErrorSendingReminderEmail:     http.StatusInternalServerError,
	ErrInvalidID:                     http.StatusBadRequest,
	ErrFilteringUsers:                http.StatusInternalServerError,
	ErrTokenNotFound:                 http.StatusNotFound,
	ErrErrorUpdatingUser:             http.StatusInternalServerError,
	ErrCredentialsNotValid:           http.StatusUnauthorized,
	ErrErrorCreatingLoan:             http.StatusInternalServerError,
	ErrLoanNotFound:                  http.StatusInternalServerError,
}

func MapErrorToStatusCode(err error) int {
	if statusCode, exists := errorStatusMap[err]; exists {
		return statusCode
	}
	return http.StatusInternalServerError
}
