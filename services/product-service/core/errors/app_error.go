package errors

type ErrorType string

const (
	Validation ErrorType = "validation"
	NotFound   ErrorType = "not_found"
	Internal   ErrorType = "internal"
	Conflict   ErrorType = "conflict"
	Forbidden  ErrorType = "forbidden"
)
const AppErrorKey = "app_error"

type AppError struct {
	Type    ErrorType
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return e.Message
}

func New(t ErrorType, msg string, err error) *AppError {
	return &AppError{Type: t, Message: msg, Err: err}
}
