package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func InitDB() {
	// 初始化数据库连接
	db := InitDBConnection()
	defer db.Close()
}

func InitDBConnection() *sql.DB {
	info := "side:Serve@123@tcp(10.30.19.15:3306)/side"
	db, err := sql.Open("mysql", info)
	if err != nil {
		log.Fatalln(err) // connect to database failed
	}

	db.SetMaxOpenConns(35)
	db.SetMaxIdleConns(35)

	if err = db.Ping(); err != nil {
		log.Fatalln(err) // database is not alive
	}

	return db
}
