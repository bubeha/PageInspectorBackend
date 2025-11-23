package main

import (
	"log"
	"net/http"

	"github.com/bubeha/PageInspectorBackend/internal/api"
	cfg "github.com/bubeha/PageInspectorBackend/internal/config"
	"github.com/bubeha/PageInspectorBackend/internal/database"
)

func main() {
	config := cfg.Load()

	db, dbErr := database.NewDb(config)

	if dbErr != nil {
		log.Fatalf("Failed to connect to database: %v", dbErr)
	}

	defer func(db *database.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Failed to close database: %v", err)
		}
	}(db)

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
