package health

import (
	"log"
	"net/http"
)

func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Alive")); err != nil {
		log.Printf("HTTP response write failed: %v", err)
	}
}
