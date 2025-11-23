package httputil

import (
	"encoding/json"
	"net/http"
)

type Responder interface {
	JSON(w http.ResponseWriter, data interface{}, status int) error
	Error(w http.ResponseWriter, msg string, status int)
}

type JSONResponder struct{}

func (r *JSONResponder) JSON(w http.ResponseWriter, data interface{}, status int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func (r *JSONResponder) Error(w http.ResponseWriter, msg string, status int) {
	http.Error(w, msg, status)
}
