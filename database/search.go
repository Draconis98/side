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

func SearchImage(db *sql.DB, imageName string) bool {
	stmt, err := db.Prepare("SELECT image_name FROM image WHERE image_name = ?")
	if err != nil {
		log.Fatalln(err)
	}

	var name string
	if err = stmt.QueryRow(imageName).Scan(&name); err != nil {
		return false
	}

	return true
}

func SearchContainer(db *sql.DB, containerName string) bool {
	stmt, err := db.Prepare("SELECT container_name FROM container WHERE container_name = ?")
	if err != nil {
		log.Fatalln(err)
	}

	var name string
	if err = stmt.QueryRow(containerName).Scan(&name); err != nil {
		return false
	}

	return true
}
