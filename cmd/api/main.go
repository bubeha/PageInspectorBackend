package main

import (
	"github.com/bubeha/PageInspectorBackend/internal/app/domain"
	cfg "github.com/bubeha/PageInspectorBackend/internal/infrastructure/config"
	"github.com/bubeha/PageInspectorBackend/internal/infrastructure/database"
	"github.com/bubeha/PageInspectorBackend/internal/interfaces/api"
	"github.com/bubeha/PageInspectorBackend/internal/repository"
	"github.com/bubeha/PageInspectorBackend/pkg/log"
	"github.com/bubeha/PageInspectorBackend/pkg/validator"
)

func main() {
	logger, lErr := log.NewZapLogger()

	if lErr != nil {
		panic(lErr)
	}

	defer func(logger *log.ZapLogger) {
		err := logger.Sync()

		if err != nil {
			log.Errorf("Failed to sync logger: %v", err)
		}
	}(logger)

	log.SetLogger(logger)

	config := initConfig()

	validator.Init()

	db := initDatabase(config)

	defer func(db *database.DB) {
		err := db.Close()

		if err != nil {
			log.Error("Failed to close database: %v", err)
		}
	}(db)

	// Create services
	domainRepo := repository.NewDomainRepository(db)
	domainService := domain.NewDomainService(domainRepo)

	server := api.NewServer(
		&api.DataLayer{DomainRepo: domainRepo},
		&api.Services{DomainService: domainService},
		&api.Infrastructure{Config: config, DB: db},
	)

	log.Info("========================================")
	log.Infof("Server starting on %s:%s", config.Server.Host, config.Server.Port)
	log.Info("========================================")

	if err := server.Run(); err != nil {
		log.Errorf("Failed to start server: %v", err)
	}
}

func initDatabase(config *cfg.Config) *database.DB {
	db, err := database.NewDb(config)

	if err != nil {
		log.Error("Failed to connect to database: %v", err)
	}

	return db
}

func initConfig() *cfg.Config {
	config, err := cfg.Load()

	if err != nil {
		log.Error("Config load error: ", err)
	}

	return config
}
