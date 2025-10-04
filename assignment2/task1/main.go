package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres password=password123 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to PostgreSQL!")

	createUsersTable(db)
	insertUser(db, "Rasul", 20)
	insertUser(db, "Azamat", 24)
	queryUsers(db)
}

func createUsersTable(db *sql.DB) {
	dropQuery := `DROP TABLE IF EXISTS users;`
	_, err := db.Exec(dropQuery)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table 'users' dropped successfully (if it existed).")

	createQuery := `
	CREATE TABLE users (
		id SERIAL PRIMARY KEY,
		name TEXT,
		age INT
	);`

	_, err = db.Exec(createQuery)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table 'users' created successfully.")
}

func insertUser(db *sql.DB, name string, age int) {
	query := "INSERT INTO users (name, age) VALUES ($1, $2)"
	_, err := db.Exec(query, name, age)
	if err != nil {
		log.Println("Failed to insert user:", err)
	} else {
		fmt.Printf("Inserted user: %s\n", name)
	}
}

func queryUsers(db *sql.DB) {
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("\nUsers:")
	for rows.Next() {
		var id int
		var name string
		var age int
		if err := rows.Scan(&id, &name, &age); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}
}
