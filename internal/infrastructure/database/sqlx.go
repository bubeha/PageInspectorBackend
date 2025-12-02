package database

import (
	"context"
	"fmt"
	"time"

	"github.com/bubeha/PageInspectorBackend/internal/infrastructure/config"
	log "github.com/bubeha/PageInspectorBackend/internal/infrastructure/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	*sqlx.DB
}

func NewDb(config *config.Config) (*DB, error) {
	dns := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", config.Database.Host, config.Database.Port, config.Database.User, config.Database.Name, config.Database.Password, config.Database.SSLMode)

	db, contErr := sqlx.Connect("postgres", dns)

	if contErr != nil {
		return nil, fmt.Errorf("failed to connect to database: %s", contErr)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to pint to database: %s", err)
	}

	log.Info("Database connected successfully (sqlx)")

	return &DB{db}, nil
}

func (db *DB) Close() error {
	return db.DB.Close()
}

func (db *DB) WithTransaction(ctx context.Context, fn func(*sqlx.Tx) error) error {
	tx, tranErr := db.BeginTxx(ctx, nil)
	if tranErr != nil {
		return fmt.Errorf("begin transaction: %w", tranErr)
	}

	defer func() {
		if p := recover(); p != nil {
			if err := tx.Rollback(); err != nil {
				return
			}
			panic(p)
		}
	}()

	if err := fn(tx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("rollback error: %v, original error: %w", rbErr, err)
		}
		return err
	}

	return tx.Commit()
}
