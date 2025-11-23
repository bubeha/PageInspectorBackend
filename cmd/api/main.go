package main

import (
	"log"
	"net/http"

	"github.com/bubeha/PageInspectorBackend/internal/api"
	cfg "github.com/bubeha/PageInspectorBackend/internal/config"
)

func main() {
	config := cfg.Load()

	router := api.Setup(&config.Server)

	err := http.ListenAndServe(
		":"+config.Server.Port,
		router,
	)

	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
