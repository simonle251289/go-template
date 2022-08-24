package app_errors

import (
	"errors"
)

type ErrorCode struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
}

type AppError struct {
	StatusCode int   `json:"-"`
	RootErr    error `json:"-"`
	ErrorCode  `json:",inline"`
}

func NewErrorResponse(root error, statusCode int, errorCode *ErrorCode) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		ErrorCode:  *errorCode,
	}
}

func NewError(root error, statusCode int, errorCode *ErrorCode) *AppError {
	if root != nil {
		return NewErrorResponse(root, statusCode, errorCode)
	}
	return NewErrorResponse(errors.New(errorCode.Message), statusCode, errorCode)
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}
	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func NewErrorCode(errorCode int, message string) *ErrorCode {
	return &ErrorCode{
		ErrorCode: errorCode,
		Message:   message,
	}
}

var (
	// Generic /Unknown error start error code in [0 -> 99]
	Generic = NewErrorCode(0, "Internal Server Error.")

	// InvalidUrl Resource error start code in [10000 -> 10999]
	InvalidUrl              = NewErrorCode(1, "Invalid URL.")
	MissingRequiredField    = NewErrorCode(2, "Missing required fields.")
	InvalidAccess           = NewErrorCode(3, "Invalid access.")
	WrongUserNameOrPassword = NewErrorCode(3, "The user name or password combination you have entered does not match our records")
	ItemNotFound            = NewErrorCode(4, "Item not found.")
)
