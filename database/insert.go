package main

import (
	"database/sql"
	"fmt"
	"log"
)

func InsertUser(db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO user(user_name) VALUES(?)")
	if err != nil {
		log.Fatalln(err)
	}

	res, err := stmt.Exec("Dolly")
	if err != nil {
		log.Fatalln(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(id)
}
