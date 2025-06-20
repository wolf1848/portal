package repository

import (
	"fmt"
	"log"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/wolf1848/gotaxi/config"
)

func InitMysql(c *config.Config) (*sql.DB, error) {

	connStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		c.Database.Mysql.User,
		c.Database.Mysql.Password,
		c.Database.Mysql.Host,
		c.Database.Mysql.Port,
		c.Database.Mysql.Dbname,
	)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer db.Close()

	// Проверяем соединение
	err = db.Ping()
	if err != nil {
		log.Fatal("Не удалось проверить соединение с базой данных:", err)
		return nil, err
	}

	log.Println("Успешное подключение к MySQL!")

	// Устанавливаем параметры пула соединений
	db.SetMaxOpenConns(10)                 // Максимальное количество открытых соединений
	db.SetMaxIdleConns(10)                 // Максимальное количество неактивных соединений
	db.SetConnMaxLifetime(5 * time.Minute) // Максимальное время жизни соединения

	return db, nil
}
