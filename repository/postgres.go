package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"gotaxi/config"
)

var Pool *pgxpool.Pool // Лучше сделать приватной, а экземпляр возвращать методом

func InitDB() error { // сейчас Init и Close это отдельные функции, я б сделал через объект с методами, можно заюзать singleton

	/*
	 А зачем исключил конфиг из гита? Пусть лежит себе структура,
	 заполнять из yaml файла, который тоже пусть себе лежит в корне проекта,
	 а данные которые нельзя палить - должны браться из переменных окружения
	*/
	connStr := config.GetDBConnectionString()

	cfg, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return fmt.Errorf("failed to parse db config: %w", err)
	}

	// Настройки пула соединений
	cfg.MaxConns = 10 // Такое тоже можно выносить в конфиг, как раз в yaml файле такое можно задавать сразу
	cfg.MinConns = 2
	cfg.MaxConnLifetime = time.Hour
	cfg.MaxConnIdleTime = time.Minute * 30

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Таймаут тоже можно в конфиг
	defer cancel()

	Pool, err = pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return fmt.Errorf("failed to create connection pool: %w", err) // я б все таки https://github.com/pkg/errors взял, но не критично
	}

	// Проверка соединения
	if err := Pool.Ping(ctx); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL") // https://github.com/sirupsen/logrus или https://github.com/uber-go/zap - тоже не критично
	return nil
}

func CloseDB() {
	if Pool != nil {
		Pool.Close()
	}
}
