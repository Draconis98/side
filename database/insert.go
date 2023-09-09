package main

import (
	"database/sql"
	"log"
	"time"
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

func InsertImage(db *sql.DB, imageName, userName string) {
	commitTime := time.Now()
	stmt, err := db.Prepare("INSERT INTO image(image_name, user_name, commit_time) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = stmt.Exec(userName, imageName, commitTime)
	if err != nil {
		log.Fatalln(err)
	}
}
