package database

import (
	"log"
	"time"
)

func InsertUser(username string) {
	if CheckUserExists(username) {
		return
	}
	stmt, err := GetDBInstance().Prepare("INSERT INTO user(user_name) VALUES(?)")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = stmt.Exec(username)
	if err != nil {
		log.Fatalln(err)
	}
}

// func InsertImage(db *sql.DB, imageName, userName string) {
// 	commitTime := time.Now()
// 	stmt, err := db.Prepare("INSERT INTO image(image_name, commit_time, user_name) VALUES(?, ?, ?)")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	_, err = stmt.Exec(imageName, commitTime, userName)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }

func InsertContainer(containerName, userName, imageName string, status, cpu, memory int, createdAt time.Time) {
	currentTime := time.Now()
	stmt, err := GetDBInstance().Prepare("INSERT INTO container(container_id, user_name, last_visit, based_image, status, core, memory, created_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Panicln(err)
	}

	_, err = stmt.Exec(containerName, userName, currentTime, imageName, status, cpu, memory, createdAt)
	if err != nil {
		log.Panicln(err)
	}
}
