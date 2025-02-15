package db

import (
	"baseToDo/dependencies"

	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"user=%s dbname=%s sslmode=disable password=%s",
		dependencies.UserDB,
		dependencies.NameDB,
		dependencies.PasswordDB,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
