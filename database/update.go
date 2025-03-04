package database

import (
	"log"
)

// func UpdateContainerStatus(db *sql.DB, containerName string, status int) bool {
// 	stmt, err := db.Prepare("UPDATE container SET status = ? WHERE container_name = ?")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	_, err = stmt.Exec(status, containerName)
// 	if err != nil {
// 		return false
// 	}

// 	return true
// }

func UpdateContainerInfo(containerName string, cpu int, memory int) bool {
	stmt, err := GetDBInstance().Prepare("UPDATE container SET core = ?, memory = ? WHERE container_id = ?")
	if err != nil {
		log.Panicln(err)
	}

	_, err = stmt.Exec(cpu, memory, containerName)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func UpdateContainerStatus(containerName string, status int) bool {
	stmt, err := GetDBInstance().Prepare("UPDATE container SET status = ? WHERE container_id = ?")
	if err != nil {
		log.Panicln(err)
	}

	_, err = stmt.Exec(status, containerName)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

// func UpdateContainerMem(db *sql.DB, containerName string, mem int) bool {
// 	stmt, err := db.Prepare("UPDATE container SET memory = ? WHERE container_name = ?")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	_, err = stmt.Exec(mem, containerName)
// 	if err != nil {
// 		return false
// 	}

// 	return true
// }

// // imgType: based_image, commit_image
// func UpdateContainerImage(db *sql.DB, containerName, imageName, imgType string) bool {
// 	stmt, err := db.Prepare("UPDATE container SET " + imgType + " = ? WHERE container_name = ?")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	_, err = stmt.Exec(imageName, containerName)
// 	if err != nil {
// 		return false
// 	}

// 	return true
// }

// // 存疑
// func UpdateLastVisit(db *sql.DB, containerName string) bool {
// 	stmt, err := db.Prepare("UPDATE container SET last_visit = ? WHERE container_name = ?")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	_, err = stmt.Exec(containerName)
// 	if err != nil {
// 		return false
// 	}

// 	return true
// }
