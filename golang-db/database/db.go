package database

import (
    "database/sql"
    "fmt"
    "log"
    "golang-db/models"
    _ "github.com/mattn/go-sqlite3"
)

func Connect() (*sql.DB, error) {
    db, err := sql.Open("sqlite3", "./test.db")
    if err != nil {
        return nil, fmt.Errorf("error opening database: %v", err)
    }
    return db, nil
}

func CreateTable(db *sql.DB) error {
    createTableSQL := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        email TEXT UNIQUE
    );`
    if _, err := db.Exec(createTableSQL); err != nil {
        return fmt.Errorf("error creating table: %v", err)
    }
    return nil
}

func InsertUser(db *sql.DB, user *models.User) error {
    insertSQL := `INSERT INTO users (name, email) VALUES (?, ?)`
    result, err := db.Exec(insertSQL, user.Name, user.Email)
    if err != nil {
        return fmt.Errorf("could not insert user: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return fmt.Errorf("could not retrieve last insert ID: %v", err)
    }
    user.ID = int(id)
    fmt.Printf("User inserted successfully with ID %d\n", user.ID)
    return nil
}

func QueryUsers(db *sql.DB) error {
    querySQL := `SELECT id, name, email FROM users`
    rows, err := db.Query(querySQL)
    if err != nil {
        return fmt.Errorf("could not query users: %v", err)
    }
    defer rows.Close()

    fmt.Println("Users:")
    for rows.Next() {
        var user models.User
        if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
            return fmt.Errorf("could not scan row: %v", err)
        }
        fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
    }

    return nil
}

func UpdateUser(db *sql.DB, user *models.User) error {
    updateSQL := `UPDATE users SET name = ?, email = ? WHERE id = ?`
    _, err := db.Exec(updateSQL, user.Name, user.Email, user.ID)
    if err != nil {
        return fmt.Errorf("could not update user: %v", err)
    }
    fmt.Println("User updated successfully")
    return nil
}

func DeleteUser(db *sql.DB, id int) error {
    deleteSQL := `DELETE FROM users WHERE id = ?`
    _, err := db.Exec(deleteSQL, id)
    if err != nil {
        return fmt.Errorf("could not delete user: %v", err)
    }
    fmt.Println("User deleted successfully")
    return nil
}
