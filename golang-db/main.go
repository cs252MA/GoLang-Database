package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/mattn/go-sqlite3"
    "golang-db/database"
    "golang-db/models"
)

func main() {
    // Connect to the SQLite database
    database, err := db.Connect()
    if err != nil {
        log.Fatalf("Error opening database: %v\n", err)
    }
    defer database.Close()

    // Create the users table
    if err := db.CreateTable(database); err != nil {
        log.Fatalf("Error creating table: %v\n", err)
    }

    fmt.Println("Users table created!")

    // Insert a new user
    user := models.User{Name: "Mark Mark", Email: "mark.mark@example.com"}
    if err := db.InsertUser(database, &user); err != nil {
        log.Printf("Error inserting user: %v\n", err)
    }

    // Query and display all users
    if err := db.QueryUsers(database); err != nil {
        log.Printf("Error querying users: %v\n", err)
    }

    // Update a user
    updatedUser := models.User{ID: 1, Name: "Alice Alice", Email: "alice.alice@example.com"}
    if err := db.UpdateUser(database, &updatedUser); err != nil {
        log.Printf("Error updating user: %v\n", err)
    }

    // Query and display all users
    if err := db.QueryUsers(database); err != nil {
        log.Printf("Error querying users: %v\n", err)
    }

    // Delete a user
    if err := db.DeleteUser(database, 1); err != nil {
        log.Printf("Error deleting user: %v\n", err)
    }

    // Query and display all users
    if err := db.QueryUsers(database); err != nil {
        log.Printf("Error querying users: %v\n", err)
    }
}
