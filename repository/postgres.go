package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wolf1848/gotaxi/config"
)

var Pool *pgxpool.Pool

func InitDB() error {
	connStr := config.GetDBConnectionString()

	cfg, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return fmt.Errorf("failed to parse db config: %w", err)
	}

	// Настройки пула соединений
	cfg.MaxConns = 10
	cfg.MinConns = 2
	cfg.MaxConnLifetime = time.Hour
	cfg.MaxConnIdleTime = time.Minute * 30

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	Pool, err = pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return fmt.Errorf("failed to create connection pool: %w", err)
	}

	// Проверка соединения
	if err := Pool.Ping(ctx); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL")
	return nil
}

func CloseDB() {
	if Pool != nil {
		Pool.Close()
	}
}
