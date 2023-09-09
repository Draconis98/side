package main

import (
	"database/sql"
	"log"
)

func DeleteUser(db *sql.DB, userName string) {
	stmt, err := db.Prepare("DELETE FROM user WHERE user_name = ?")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = stmt.Exec(userName)
	if err != nil {
		log.Fatalln(err)
	}
}

func DeleteContainer(db *sql.DB, containerName string) bool {
	stmt, err := db.Prepare("DELETE FROM container WHERE container_name = ?")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = stmt.Exec(containerName)
	if err != nil {
		return false
	}

	return true
}
