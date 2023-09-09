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

func GetResourceInfoByContainerId(containerId string) (int, int) {
	if flag := CheckContainerExists(containerId); !flag {
		log.Printf("container %v not exist", containerId)
		return -1, -1
	}

	stmt, err := GetDBInstance().Prepare("SELECT core, memory FROM container WHERE container_id = ?")
	if err != nil {
		log.Panicln(err)
	}

	var cpu, memory int
	if err = stmt.QueryRow(containerId).Scan(&cpu, &memory); err != nil {
		log.Panicln(err)
	}

	return cpu, memory
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

func CheckContainerExists(containerId string) bool {
	stmt, err := GetDBInstance().Prepare("SELECT container_id FROM container WHERE container_id = ?")
	if err != nil {
		log.Panicln(err)
	}

	var name string
	if err = stmt.QueryRow(containerId).Scan(&name); err != nil {
		return false
	}

	return true
}

// Get the containers by username.
// If the user doesn't exist, return nil
// If the user have no containers, return an empty list.
func GetContainersByUser(username string) []utils.Container {
	if flag := CheckUserExists(username); !flag {
		log.Printf("User %v not exist", username)
		return nil
	}

	stmt, err := GetDBInstance().Prepare("SELECT container_id, core, memory, status, created_at FROM container WHERE user_name = ?")
	if err != nil {
		log.Panicln(err)
	}

	rows, err := stmt.Query(username)
	if err != nil {
		log.Panicln(err)
	}

	containers := make([]utils.Container, 0)
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
