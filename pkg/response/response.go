package response

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/bubeha/PageInspectorBackend/pkg/validator"
)

func JSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Error(w http.ResponseWriter, msg string, status int) {
	http.Error(w, msg, status)
}

func JsonError(w http.ResponseWriter, err error, status int) {
	w.Header().Set("Content-Type", "application/json")

	var validationErr validator.ValidationError

	if errors.As(err, &validationErr) {
		w.WriteHeader(status)
		if encodeErr := json.NewEncoder(w).Encode(validationErr); encodeErr != nil {
			w.Header().Set("Content-Type", "text/plain")
			http.Error(w, http.StatusText(status), status)
		}
		return
	}

	w.WriteHeader(status)
	if jErr := json.NewEncoder(w).Encode(err.Error()); jErr != nil {
		http.Error(w, http.StatusText(status), status)
	}
}
