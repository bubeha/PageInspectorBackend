package response

import (
	"encoding/json"
	"net/http"
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
