package errors

import (
	"fmt"
	"net/http"
)

// ErrorType represents the type of error
type ErrorType string

const (
	// ErrorTypeInternal represents internal server errors
	ErrorTypeInternal ErrorType = "INTERNAL"

	// ErrorTypeBadRequest represents client-side errors
	ErrorTypeBadRequest ErrorType = "BAD_REQUEST"

	// ErrorTypeNotFound represents resource not found errors
	ErrorTypeNotFound ErrorType = "NOT_FOUND"

	// ErrorTypeUnauthorized represents authentication errors
	ErrorTypeUnauthorized ErrorType = "UNAUTHORIZED"

	// ErrorTypeForbidden represents authorization errors
	ErrorTypeForbidden ErrorType = "FORBIDDEN"

	// ErrorTypeTimeout represents timeout errors
	ErrorTypeTimeout ErrorType = "TIMEOUT"

	// ErrorTypeConflict represents resource conflict errors
	ErrorTypeConflict ErrorType = "CONFLICT"
)

// AppError represents an application error
type AppError struct {
	Type    ErrorType
	Message string
	Cause   error
	Code    int
}

// Error returns the error message
func (e AppError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %s (caused by: %s)", e.Type, e.Message, e.Cause.Error())
	}
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

// StatusCode returns the HTTP status code for the error
func (e AppError) StatusCode() int {
	if e.Code != 0 {
		return e.Code
	}

	switch e.Type {
	case ErrorTypeInternal:
		return http.StatusInternalServerError
	case ErrorTypeBadRequest:
		return http.StatusBadRequest
	case ErrorTypeNotFound:
		return http.StatusNotFound
	case ErrorTypeUnauthorized:
		return http.StatusUnauthorized
	case ErrorTypeForbidden:
		return http.StatusForbidden
	case ErrorTypeTimeout:
		return http.StatusGatewayTimeout
	case ErrorTypeConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

// NewInternalError creates a new internal server error
func NewInternalError(message string, cause error) AppError {
	return AppError{
		Type:    ErrorTypeInternal,
		Message: message,
		Cause:   cause,
	}
}

// NewBadRequestError creates a new bad request error
func NewBadRequestError(message string, cause error) AppError {
	return AppError{
		Type:    ErrorTypeBadRequest,
		Message: message,
		Cause:   cause,
	}
}

// NewNotFoundError creates a new not found error
func NewNotFoundError(message string, cause error) AppError {
	return AppError{
		Type:    ErrorTypeNotFound,
		Message: message,
		Cause:   cause,
	}
}

// NewUnauthorizedError creates a new unauthorized error
func NewUnauthorizedError(message string, cause error) AppError {
	return AppError{
		Type:    ErrorTypeUnauthorized,
		Message: message,
		Cause:   cause,
	}
}

// NewForbiddenError creates a new forbidden error
func NewForbiddenError(message string, cause error) AppError {
	return AppError{
		Type:    ErrorTypeForbidden,
		Message: message,
		Cause:   cause,
	}
}

// NewTimeoutError creates a new timeout error
func NewTimeoutError(message string, cause error) AppError {
	return AppError{
		Type:    ErrorTypeTimeout,
		Message: message,
		Cause:   cause,
	}
}

// NewConflictError creates a new conflict error
func NewConflictError(message string, cause error) AppError {
	return AppError{
		Type:    ErrorTypeConflict,
		Message: message,
		Cause:   cause,
	}
}
