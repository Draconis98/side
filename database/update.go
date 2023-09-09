package main

import (
	"database/sql"
	"log"
)

func UpdateContainerStatus(db *sql.DB, containerName string, status int) bool {
	stmt, err := db.Prepare("UPDATE container SET status = ? WHERE container_name = ?")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = stmt.Exec(status, containerName)
	if err != nil {
		return false
	}

	return true
}
