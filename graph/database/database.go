package database

import (
	"database/sql"
	"log"
	"os"

	// DB
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func InitDB() {
	cnx := os.Getenv("UNAME") + ":" + os.Getenv("PASS")
	db, err := sql.Open("mysql", cnx+"@("+os.Getenv("SERVER")+":3306)/stacke")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	Db = db
}
