package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//var DB *sql.DB

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

func UpdateCoilStatus(HEX string, status int, Client string) {
	if status == '1' {
		coilGoingToClient(HEX, Client)
	} else if status == '0' {
		coilReturning(HEX)
	} else if status == '2' {
		CoilToMan(HEX)
	} else {
		log.Fatalf("Invalid Status!")
	}
}

func coilGoingToClient(HEX string, Client string) {
	InitDB()
	stmt, err := DB.Prepare("UPDATE coil SET Status = 1, Client = ? WHERE HEX = ?;")
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

func coilReturning(HEX string) {
	InitDB()
	stmt, err := DB.Prepare("UPDATE coil SET Status = 0, Cycles = Cycles + 1 WHERE HEX = ?;")
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
	stmt, err := DB.Prepare("UPDATE coils SET Status = 2, Cycles = Cycles + 1, Client = NULL WHERE HEX = ?;")
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

func UpdateRackStatus(HEX string, status int, Client string) {
	if status == '1' {
		rackGoingToClient(HEX, Client)
	} else if status == '0' {
		rackReturning(HEX)
	} else if status == '2' {
		rackToMan(HEX)
	} else {
		log.Fatalf("Invalid Status!")
	}
}

func rackGoingToClient(HEX string, Client string) {
	InitDB()
	stmt, err := DB.Prepare("UPDATE rack SET Status = 1, Client = ? WHERE HEX = ?;")
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
		log.Printf("No rack was found or updated for HEX: %s", HEX)
	} else {
		log.Printf("Successfully updated the rack with HEX: %s", HEX)
	}
}

func rackReturning(HEX string) {
	InitDB()
	stmt, err := DB.Prepare("UPDATE rack SET Status = 0, Cycles = Cycles + 1 WHERE HEX = ?;")
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
		log.Printf("No rack was found or updated for HEX: %s", HEX)
	} else {
		log.Printf("Successfully updated the rack with HEX: %s", HEX)
	}
}

func rackToMan(HEX string) {
	InitDB()
	stmt, err := DB.Prepare("UPDATE rack SET Status = 2, Cycles = Cycles + 1, Client = NULL WHERE HEX = ?;")
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
		log.Printf("No rack was found or updated for HEX: %s", HEX)
	} else {
		log.Printf("Successfully updated the rack with HEX: %s", HEX)
	}
}
