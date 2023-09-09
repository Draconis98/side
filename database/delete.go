package database

import (
	"log"
)

func DeleteContainer(containerId string) bool {
	if flag := CheckContainerExists(containerId); !flag {
		log.Printf("Container %v not exist", containerId)
		return false
	}

	stmt, err := GetDBInstance().Prepare("DELETE FROM container WHERE container_id = ?")
	if err != nil {
		log.Panicln(err)
	}

	_, err = stmt.Exec(containerId)
	if err != nil {
		log.Panicln(err)
	}

	return true
}
