package main

import (
	"database/sql"
	"log"
)

func DeleteUser(db *sql.DB, userName string) {
	stmt, err := db.Prepare("DELETE FROM user WHERE user_name = ?")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = stmt.Exec(userName)
	if err != nil {
		log.Fatalln(err)
	}
}
