package rest_err

import (
	"net/http"
)

type RestErr struct {
	Message string`json:"message"`
	Err string`json:"error"`
	Code int`json:"code"`
	Causes []Causes`json:"causes,omitempty"`
}

type Causes struct {
	Field string`json:"field"`
	Message string`json:"message"`
}

func NewRestErr(message string, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err: err,
		Code: code,
		Causes: causes,
	}
}

// 400 - Bad Request - General Error
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "bad request",
		Code: http.StatusBadRequest,
	}
}

// 400 - Bad Request - Validation Error
func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err: "bad request",
		Code: http.StatusBadRequest,
		Causes: causes,
	}
}

// 500 - Internal Server Error - General Error
func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "internal server error",
		Code: http.StatusInternalServerError,
	}
}

// 500 - Internal Server Error - Validation Error
func NewInternalServeValidationrError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err: "internal server error",
		Code: http.StatusInternalServerError,
		Causes: causes,
	}
}

// 404 - Not Found - General Error
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "not found",
		Code: http.StatusNotFound,
	}
}

// 404 - Not Found - Validation Error
func NewNotFoundValidationrError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err: "not found",
		Code: http.StatusNotFound,
		Causes: causes,
	}
}

// 401 - Unauthorized - General Error
func NewUnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "unauthorized",
		Code: http.StatusUnauthorized,
	}
}

// 401 - Unauthorized - Validation Error
func NewUnauthorizedValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err: "unauthorized",
		Code: http.StatusUnauthorized,
		Causes: causes,
	}
}

// 403 - Forbidden - General Error
func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err: "forbidden",
		Code: http.StatusForbidden,
	}
}

// 403 - Forbidden - Validation Error
func NewForbiddenValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err: "forbidden",
		Code: http.StatusForbidden,
		Causes: causes,
	}
}

func (r *RestErr) Error() string {
	return r.Message
}