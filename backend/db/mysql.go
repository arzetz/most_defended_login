package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	dsn := "root:8499k8499k@tcp(localhost:3306)/auth"
	database, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения: %v", err)
	}
	if err := database.Ping(); err != nil {
		return nil, fmt.Errorf("ошибка пинга: %v", err)
	} else {
		fmt.Printf("ЧУВААААК")
	}
	return database, nil
}
