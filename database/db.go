package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go_mysql")
	if err != nil {
		panic(err)
	}

	return db
}
