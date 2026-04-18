package errors

import "errors"

var Is = errors.Is
var Join = errors.Join

var ErrUserNotFound = errors.New("user not found")
