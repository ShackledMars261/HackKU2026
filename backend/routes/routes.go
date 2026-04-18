package routes

import (
	"encoding/json"
	"main/errors"
	"net/http"
)

func writeJSON(w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, err error) {
	if clientErr, ok := errors.AsType[*errors.ClientError](err); ok {
		writeJSON(w, clientErr.Code, clientErr)
		return
	}

	writeJSON(w, http.StatusInternalServerError, errors.ErrUnknown)
}

func readBody(r *http.Request, data any) error {
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		return errors.ErrInvalidBody
	}

	return nil
}
