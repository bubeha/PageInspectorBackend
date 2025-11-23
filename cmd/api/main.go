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

	log.Printf("========================================")
	log.Printf("🤖 Server started successfully!")
	log.Printf("🌐 URL: http://%s:%s", config.Server.Host, config.Server.Port)
	log.Printf("========================================")

	err := http.ListenAndServe(
		config.Server.Host+":"+config.Server.Port,
		router,
	)

	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
