package main

import (
	"database/sql"
	"log"
)

func SearchUser(db *sql.DB, userName string) bool {
	stmt, err := db.Prepare("SELECT user_name FROM user WHERE user_name = ?")
	if err != nil {
		log.Fatalln(err)
	}

	var name string
	if err = stmt.QueryRow(userName).Scan(&name); err != nil {
		return false
	}

	return true
}
