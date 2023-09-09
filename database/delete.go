package database

import (
	"log"
)

func DeleteContainer(containerName string) bool {
	stmt, err := GetDBInstance().Prepare("DELETE FROM container WHERE container_id = ?")
	if err != nil {
		log.Panicln(err)
	}

	_, err = stmt.Exec(containerName)
	if err != nil {
		return false
	}

	return true
}
