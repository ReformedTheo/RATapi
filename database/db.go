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
	if err != nil {
		log.Panic(err)
	}

	defer DB.Close()

	if err = DB.Ping(); err != nil {
		log.Panic(err)
	}
}

func UpdateCoilStatus(HEX string, state int, Client string) {
	InitDB()
	if state == '1' {
		CoilGoingToClient(HEX, Client)
	} else if state == '0' {
		CoilReturning(HEX)
	}
}

func CoilGoingToClient(HEX string, Client string) {
	InitDB()
	stmt, err := DB.Prepare("UPDATE Coil SET Status = 1, Client = ? WHERE HEX = ?;")
	if err != nil {
		log.Fatalf("Could not prepare sql statement: %v", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(Client, HEX)
	if err != nil {
		log.Fatalf("Could not execute sql statement: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error when checking rows affected: %v", err)
	}

	if rowsAffected == 0 {
		log.Printf("No coil was found or updated for HEX: %s", HEX)
	} else {
		log.Printf("Successfully updated the coil with HEX: %s", HEX)
	}
}

func CoilReturning(HEX string) {
	InitDB()
	stmt, err := DB.Prepare("UPDATE Coil SET Status = 0, Cycles = Cycles + 1 WHERE HEX = ?;")
	if err != nil {
		log.Fatalf("Could not prepare sql statement: %v", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(HEX)
	if err != nil {
		log.Fatalf("Could not execute sql statement: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error when checking rows affected: %v", err)
	}

	if rowsAffected == 0 {
		log.Printf("No coil was found or updated for HEX: %s", HEX)
	} else {
		log.Printf("Successfully updated the coil with HEX: %s", HEX)
	}
}

func CoilToMan(HEX string) {
	InitDB()
	stmt, err := DB.Prepare("UPDATE Coil SET Status = 2, Cycles = Cycles + 1, Client = NULL WHERE HEX = ?;")
	if err != nil {
		log.Fatalf("Could not prepare sql statement: %v", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(HEX)
	if err != nil {
		log.Fatalf("Could not execute sql statement: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error when checking rows affected: %v", err)
	}

	if rowsAffected == 0 {
		log.Printf("No coil was found or updated for HEX: %s", HEX)
	} else {
		log.Printf("Successfully updated the coil with HEX: %s", HEX)
	}
}
