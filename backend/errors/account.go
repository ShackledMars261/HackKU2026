package errors

import "errors"

var ErrUsernameAlreadyExists = errors.New("username already exists")
var ErrInvalidUsernameOrPassword = errors.New("invalid username or password")
