package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "root:1234@tcp(localhost:3306)/rfid_tracking")
	defer DB.Close()
	if err != nil {
		log.Panic(err)
	}

	if err = DB.Ping(); err != nil {
		log.Panic(err)
	}
}

func UpdateCoilStatus(HEX string, state int) {
	InitDB()
	if state == '1' {
		CoilArriving(HEX, state)
	}
}

func CoilArriving(HEX string, satate int) {

}
