package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/wolf1848/gotaxi/config"
)

func InitDB(c *config.Config) (*pgxpool.Pool, error) {

	var pool *pgxpool.Pool

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Postgres.Host,
		c.Database.Postgres.Port,
		c.Database.Postgres.User,
		c.Database.Postgres.Password,
		c.Database.Postgres.Dbname,
		c.Database.Postgres.Ssl,
	)

	cfg, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatalf("Failed to parse db config: %s", err)
		return nil, err
	}

	// Настройки пула соединений
	cfg.MaxConns = 10
	cfg.MinConns = 2
	cfg.MaxConnLifetime = time.Hour
	cfg.MaxConnIdleTime = time.Minute * 30

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err = pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		log.Fatalf("Failed to create connection pool: %s", err)
		return nil, err
	}

	// Проверка соединения
	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Failed to ping database: %s", err)
		return nil, err
	}

	log.Println("Successfully connected to PostgreSQL")

	return pool, nil
}

func Close(p *pgxpool.Pool) {
	if p != nil {
		p.Close()
	}
}
