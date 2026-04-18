package errors

import (
	"errors"
)

var Is = errors.Is
var Join = errors.Join

func AsType[E error](err error) (E, bool) {
	return errors.AsType[E](err)
}

type ClientError struct {
	Type    string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"-"`
}

func (err *ClientError) Error() string {
	return err.Message
}

func New(tag string, message string, code int) error {
	return &ClientError{
		Type:    tag,
		Message: message,
		Code:    code,
	}
}
