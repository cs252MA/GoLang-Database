package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Connect to the SQLite database
    db, err := sql.Open("sqlite3", "./test.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Create the users table
    createTableSQL := `CREATE TABLE IF NOT EXISTS users (
        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "name" TEXT,
        "email" TEXT UNIQUE
    );`
    _, err = db.Exec(createTableSQL)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Users table created!")
// CRUD User Operations
    // Insert a new user
    insertUser(db, "Mark Mark", "mark.mark@example.com")

    // Query and display all users
    queryUsers(db)

    // Update a user
    updateUser(db, 1, "Alice Alice", "alice.alice@example.com")

    // Query and display all users
    queryUsers(db)

    // Delete a user
    deleteUser(db, 1)

    // Query and display all users
    queryUsers(db)
}

func insertUser(db *sql.DB, name string, email string) {
   
}

func queryUsers(db *sql.DB) {
   
}

func updateUser(db *sql.DB, id int, name string, email string) {
    
}

func deleteUser(db *sql.DB, id int) {
   
}
