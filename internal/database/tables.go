package database

import "fmt"

func InitTables() {
	createUserTable()
}

func createUserTable() {
	query := `CREATE TABLE IF NOT EXISTS users(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_name TEXT NOT NULL UNIQUE,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	role TEXT NOT NULL,
	created_at DATETIME NOT NULL,
	updated_at DATETIME NOT NULL
	)
	`
	_, err := DB.Exec(query)

	if err != nil {
		fmt.Println(err)
		panic("users table creation failed.")
	}
	fmt.Println("users table created...")
}
