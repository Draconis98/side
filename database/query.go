package database

import (
	"log"
	"main/utils"
)

func CheckUserExists(username string) bool {
	stmt, err := GetDBInstance().Prepare("SELECT user_name FROM user WHERE user_name = ?")
	if err != nil {
		log.Panicln(err)
	}

	var name string
	if err = stmt.QueryRow(username).Scan(&name); err != nil {
		return false
	}

	return true
}

func GetResourceLimitByUser(username string) (int, int) {
	if flag := CheckUserExists(username); !flag {
		log.Printf("user %v not exist", username)
		return -1, -1
	}

	stmt, err := GetDBInstance().Prepare("SELECT cpu_limit, mem_limit FROM user WHERE user_name = ?")
	if err != nil {
		log.Panicln(err)
	}

	var cpuLimit, memLimit int
	if err = stmt.QueryRow(username).Scan(&cpuLimit, &memLimit); err != nil {
		log.Panicln(err)
	}

	return cpuLimit, memLimit
}

// func SearchImage(db *sql.DB, imageName string) bool {
// 	stmt, err := db.Prepare("SELECT image_name FROM image WHERE image_name = ?")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	var name string
// 	if err = stmt.QueryRow(imageName).Scan(&name); err != nil {
// 		return false
// 	}

// 	return true
// }

// func SearchContainer(db *sql.DB, containerName string) bool {
// 	stmt, err := db.Prepare("SELECT container_name FROM container WHERE container_name = ?")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	var name string
// 	if err = stmt.QueryRow(containerName).Scan(&name); err != nil {
// 		return false
// 	}

// 	return true
// }

func GetContainersByUser(username string) []utils.Container {
	stmt, err := GetDBInstance().Prepare("SELECT container_id, core, memory, status, created_at FROM container WHERE user_name = ?")
	if err != nil {
		log.Panicln(err)
	}

	rows, err := stmt.Query(username)
	if err != nil {
		log.Panicln(err)
	}

	var containers []utils.Container
	for rows.Next() {
		var container utils.Container
		if err := rows.Scan(
			&container.ContainerId,
			&container.Core,
			&container.Memory,
			&container.Status,
			&container.CreateAt); err != nil {
			log.Panicln(err)
		}
		containers = append(containers, container)
	}

	return containers
}
