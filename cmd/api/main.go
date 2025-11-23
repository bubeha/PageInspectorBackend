package main

import (
	"log"

	"github.com/bubeha/PageInspectorBackend/internal/app/domain"
	cfg "github.com/bubeha/PageInspectorBackend/internal/config"
	"github.com/bubeha/PageInspectorBackend/internal/database"
	"github.com/bubeha/PageInspectorBackend/internal/interfaces/api"
	"github.com/bubeha/PageInspectorBackend/internal/repository"
	"github.com/bubeha/PageInspectorBackend/pkg/httputil"
)

func main() {
	config, err := cfg.Load()

	if err != nil {
		log.Fatal("Config load error: ", err)
	}

	db, err := database.NewDb(config)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer func(db *database.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Failed to close database: %v", err)
		}
	}(db)

	// Create services
	domainRepo := repository.NewDomainRepository(db)
	domainService := domain.NewDomainService(domainRepo)

	server := api.NewServer(
		&api.DataLayer{DomainRepo: domainRepo},
		&api.Services{DomainService: domainService},
		&api.Infrastructure{Config: config, DB: db, Responser: &httputil.JSONResponder{}},
	)

	log.Printf("========================================")
	log.Printf("Server starting on %s:%s", config.Server.Host, config.Server.Port)
	log.Printf("========================================")

	if err := server.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
