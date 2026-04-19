package errors

import "net/http"

var ErrInvalidDurationFormat = New("invalid_duration_format", "Unable to parse duration string", http.StatusBadRequest)
