package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	dsn := "host=localhost user=postgres password=password123 dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	fmt.Println("Successfully connected to PostgreSQL with GORM!")

	db.AutoMigrate(&User{})
	fmt.Println("Table 'users' migrated successfully.")

	fmt.Println("Inserting users with GORM...")
	db.Create(&User{Name: "Charlie", Age: 35})
	db.Create(&User{Name: "Diana", Age: 28})
	fmt.Println("Users inserted.")

	var users []User
	db.Find(&users)

	fmt.Println("\nUsers (retrieved with GORM):")
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
	}
}
