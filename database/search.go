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

type Container struct {
	containerName string
	basedImage    string
	status        int
	cpu           int
	memory        int
}

func SearchContainerByUser(db *sql.DB, userName string) []Container {
	stmt, err := db.Prepare("SELECT container_name, based_image, status, cpu, memory FROM container WHERE user_name = ?")
	if err != nil {
		log.Fatalln(err)
	}

	rows, err := stmt.Query(userName)
	if err != nil {
		log.Fatalln(err)
	}

	var containers []Container
	for rows.Next() {
		var container Container
		if err := rows.Scan(&container.containerName, &container.basedImage,
			&container.status, &container.cpu, &container.memory); err != nil {
			log.Fatalln(err)
		}

		containers = append(containers, container)
	}

	return containers
}
