package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	// 初始化数据库连接
	db := InitDBConnection()
	defer db.Close()

	InsertUser(db)

	fmt.Println("InitDBConnection success")
}

func InitDBConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:Agileserve@123@tcp(10.30.19.15:3306)/side")
	if err != nil {
		log.Fatalln(err) // connect to database failed
	}

	return db
}
