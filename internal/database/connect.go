package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDatabase() {
	var err error
	DB, err = sql.Open("sqlite3", "blog.db")
	if err != nil {
		panic("database initialization failed.")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(15)

	InitTables()
}
