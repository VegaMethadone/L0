package connection

import (
	"L0/internal/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func getConnectionStr() string {
	conf, err := config.GetConfig()
	if err != nil {
		log.Printf("Config is damaged: %v", err)
	}
	str := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s", conf.Postgres.Username, conf.Postgres.Password, conf.Postgres.DatabaseName, conf.Postgres.Sslmode, conf.Postgres.Host)
	return str
}

func DB() (*sql.DB, error) {
	// Открываем соединение с базой данных PostgreSQL, используя строку подключения connStr.
	db, err := sql.Open("postgres", getConnectionStr())
	if err != nil {
		// Если произошла ошибка при открытии соединения, записываем ошибку в журнал и возвращаем nil и ошибку.
		log.Fatal(err)
		return nil, err
	}

	// Возвращаем указатель на объект *sql.DB и nil ошибки, если соединение успешно открыто.
	return db, nil
}
