package health

import (
	"net/http"
)

func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Alive!"))
}
