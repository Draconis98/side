package main

import (
	"database/sql"
	"fmt"
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
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}

	commitTime := time.Now().In(loc)
	stmt, err := db.Prepare("INSERT INTO image(image_name, commit_time, user_name) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = stmt.Exec(imageName, commitTime, userName)
	if err != nil {
		log.Fatalln(err)
	}
}
