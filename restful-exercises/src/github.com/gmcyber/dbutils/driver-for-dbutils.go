package main

import (
    "database/sql"
    "log"
    "github.com/mattn/go-sqlite3"
    "github.com/gmcyber/dbutils"
)

func main() {
    // Connect to Database
    db, err := sql.Open("sqlite3", "./railapi.db")
    if err != nil {
        log.Println("Driver creation failed!")
    }
    // Create tables
    dbutils.Initialize(db)
}
