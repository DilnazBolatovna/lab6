package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open SQLite database file
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create users table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY,
			first_name TEXT,
			last_name TEXT,
			email TEXT UNIQUE,
			password TEXT
		)
	`)
	if err != nil {
		panic(err)
	}

	// Create orders table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS orders (
			id INTEGER PRIMARY KEY,
			user_id INTEGER REFERENCES users(id),
			date DATE,
			status TEXT
		)
	`)
	if err != nil {
		panic(err)
	}

	// Create products table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS products (
			id INTEGER PRIMARY KEY,
			name TEXT,
			description TEXT,
			price REAL,
			quantity INTEGER
		)
	`)
	if err != nil {
		panic(err)
	}

	fmt.Println("Database created successfully")
}
