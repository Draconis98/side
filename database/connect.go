package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Init the database connection. Call this funcion in main.go.
func InitDBConnection() {
	info := ""
	db, err := sql.Open("mysql", info)

	if err != nil {
		log.Fatalln(err) // connect to database failed
	}

	db.SetMaxOpenConns(35)
	db.SetMaxIdleConns(35)

	if err = db.Ping(); err != nil {
		log.Fatalln(err) // database is not alive
	}
	DB = db
}

// Get the database instance.
func GetDBInstance() *sql.DB {
	return DB
}

// Close the database connection.
func CloseDBConnection() {
	DB.Close()
}
