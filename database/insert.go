package main

import (
	"database/sql"
	"log"
)

func InsertUser(db *sql.DB, userName string) {
	stmt, err := db.Prepare("INSERT INTO user(user_name) VALUES(?)")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = stmt.Exec(userName)
	if err != nil {
		log.Fatalln(err)
	}
}

func InsertImage(db *sql.DB, imageName, userName string, commitTime int64) {
	stmt, err := db.Prepare("INSERT INTO image(image_name, user_name, commitTime) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = stmt.Exec(userName, imageName, commitTime)
	if err != nil {
		log.Fatalln(err)
	}
}
