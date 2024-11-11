package db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/denisenkom/go-mssqldb" // MSSQL driver
)

var db *sql.DB

// InitDB initializes the connection to the MSSQL database
func InitDB() {
	var err error
	connString := "server=your-server;user id=your-username;password=your-password;database=your-db-name"

	db, err = sql.Open("mssql", connString)
	if err != nil {
		log.Fatalf("Error opening MSSQL database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	fmt.Println("Connected to MSSQL Database")
}

// GetDB returns the active database connection
func GetDB() *sql.DB {
	return db
}
