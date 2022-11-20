package wraperr

// New create error instance
func New(statusCode int, code int, message string, previous error) *Error {
	return &Error{
		StatusCode: statusCode,
		Code:       code,
		Message:    message,
		Previous:   previous,
	}
}

// Error error contain http status and code
type Error struct {
	StatusCode int
	Code       int
	Message    string
	Previous   error
}

func (e *Error) Error() string {
	return e.Message
}
