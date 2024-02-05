package database

import (
	"RATapi/models"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql", "root:1234@tcp(localhost:3306)/rfid_tracking")
	if err != nil {
		log.Panic(err)
	}

	if err = DB.Ping(); err != nil {
		log.Panic(err)
	}
}

func UpdateTransportStatus(hex string, coils []models.Coil, client string, status int) {
	switch status {
	case 0:
		transportReturning(hex, client, coils)
	case 1:
		transportSent(hex, client, coils)
	}
}

func CoilMaintence(coil_hex string) {

	//Function to set coil to maintence mode

	now := time.Now().Format("02/01/2006")
	stmtCoil, err := DB.Prepare("UPDATE coils SET status = 2, client = NULL, last_changed = ?, last_maintence = ?, rack_hex = NULL WHERE hex = ?;")
	if err != nil {
		log.Fatalf("Could not prepare sql statement: %v", err)
	}
	defer stmtCoil.Close()

	res, err := stmtCoil.Exec(now, now, coil_hex)
	if err != nil {
		log.Fatalf("Could not execute sql statement: %v", err)
	}

	rowsAffectedCoil, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error when checking rows affected: %v", err)
	}

	if rowsAffectedCoil == 0 {
		log.Printf("No Coil was found or updated for HEX: %s", coil_hex)
	} else {
		log.Printf("Successfully updated the coil with hex: %s", coil_hex)
	}
}

func transportSent(hex string, client string, coils []models.Coil) {

	now := time.Now().Format("02/01/2006")

	//Execute the update in DB for the Racks leaving for clients

	stmtRack, err := DB.Prepare("UPDATE racks SET status = 1, client = ?, last_changed = ?, cycles = cycles + 1 WHERE hex = ?;")
	if err != nil {
		log.Fatalf("Could not prepare sql statement: %v", err)
	}
	defer stmtRack.Close()

	res, err := stmtRack.Exec(client, now, hex)
	if err != nil {
		log.Fatalf("Could not execute sql statement: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error when checking rows affected: %v", err)
	}

	if rowsAffected == 0 {
		log.Printf("No rack was found or updated for HEX: %s", hex)
	} else {
		log.Printf("Successfully updated the coil with hex: %s", hex)
	}

	//Execute the update in DB for the Coils leaving for clients

	for i := 0; i < len(coils); i++ {
		stmtCoil, err := DB.Prepare("UPDATE coils SET status = 1, client = ?, last_changed = ?, rack_hex = ?,cycles = cycles + 1 WHERE hex = ?;")
		if err != nil {
			log.Fatalf("Could not prepare sql statement: %v", err)
		}
		defer stmtCoil.Close()

		res, err = stmtCoil.Exec(client, now, hex, coils[i].HEX)
		if err != nil {
			log.Fatalf("Could not execute sql statement: %v", err)
		}

		rowsAffectedCoil, err := res.RowsAffected()
		if err != nil {
			log.Fatalf("Error when checking rows affected: %v", err)
		}

		if rowsAffectedCoil == 0 {
			log.Printf("No Coil was found or updated for HEX: %s", hex)
		} else {
			log.Printf("Successfully updated the coil with hex: %s", hex)
		}
	}
}
func transportReturning(hex string, client string, coils []models.Coil) {

	now := time.Now().Format("02/01/2006")

	//Execute the update in DB for the Racks returning to production

	stmtRack, err := DB.Prepare("UPDATE racks SET status = 0, client = NULL, last_changed = ?, WHERE hex = ?;")
	if err != nil {
		log.Fatalf("Could not prepare sql statement: %v", err)
	}
	defer stmtRack.Close()

	res, err := stmtRack.Exec(now, hex)
	if err != nil {
		log.Fatalf("Could not execute sql statement: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error when checking rows affected: %v", err)
	}

	if rowsAffected == 0 {
		log.Printf("No rack was found or updated for HEX: %s", hex)
	} else {
		log.Printf("Successfully updated the coil with hex: %s", hex)
	}

	//Execute the update in DB for the Coils returning to production

	for i := 0; i < len(coils); i++ {
		stmtCoil, err := DB.Prepare("UPDATE coils SET status = 0, client = NULL, last_changed = ?, rack_hex = NULL, WHERE hex = ?;")
		if err != nil {
			log.Fatalf("Could not prepare sql statement: %v", err)
		}
		defer stmtCoil.Close()

		res, err = stmtCoil.Exec(client, now, hex, coils[i].HEX)
		if err != nil {
			log.Fatalf("Could not execute sql statement: %v", err)
		}

		rowsAffectedCoil, err := res.RowsAffected()
		if err != nil {
			log.Fatalf("Error when checking rows affected: %v", err)
		}

		if rowsAffectedCoil == 0 {
			log.Printf("No Coil was found or updated for HEX: %s", hex)
		} else {
			log.Printf("Successfully updated the coil with hex: %s", hex)
		}
	}
}
