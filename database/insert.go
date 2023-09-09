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
	stmt, err := db.Prepare("INSERT INTO image(image_name, commit_time, user_name) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = stmt.Exec(imageName, commitTime, userName)
	if err != nil {
		log.Fatalln(err)
	}
}

func InsertContainer(db *sql.DB, containerName, userName, imageName string, status, cpu, memory int) {
	currentTime := time.Now()
	stmt, err := db.Prepare("INSERT INTO container(container_name, user_name, last_visit, based_image, status, cpu, memory) VALUES(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = stmt.Exec(containerName, userName, currentTime, imageName, status, cpu, memory)
	if err != nil {
		log.Fatalln(err)
	}
}
