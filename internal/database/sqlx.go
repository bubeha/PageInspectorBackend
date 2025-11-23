package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bubeha/PageInspectorBackend/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DB struct {
	*sqlx.DB
}

func NewDb(config *config.Config) (*DB, error) {
	dns := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", config.DB.Host, config.DB.Port, config.DB.User, config.DB.Name, config.DB.Password, "disable")

	fmt.Println(123, dns)

	db, contErr := sqlx.Connect("postgres", dns)

	if contErr != nil {
		return nil, fmt.Errorf("failed to connect to database: %s", contErr)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to pint to database: %s", err)
	}

	log.Printf("Database connected successfully (sqlx)")

	return &DB{db}, nil
}

func (db *DB) Close() error {
	return db.DB.Close()
}

// WithTransaction выполняет операцию в транзакции
func (db *DB) WithTransaction(ctx context.Context, fn func(*sqlx.Tx) error) error {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
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
