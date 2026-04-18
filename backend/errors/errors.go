package errors

import "errors"

var Is = errors.Is
var Join = errors.Join

var ErrNotFound = errors.New("Not found")
