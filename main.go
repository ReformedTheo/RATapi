package main

import (
	"RATapi/database"
	"RATapi/handlers"
	"log"
	"net/http"
)

func main() {
	database.InitDB()

	http.HandleFunc("/update_status", handlers.UpdateStatusHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
