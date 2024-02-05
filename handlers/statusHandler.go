package handlers

import (
	"RATapi/database"
	"RATapi/models"
	"encoding/json"
	"net/http"
)

func UpdateStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	var requestData models.RequestData
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if requestData.Status == 2 {
		database.CoilMaintence(requestData.Coil_HEX)
	} else {
		database.UpdateTransportStatus(requestData.HEX, requestData.Coils, requestData.Client, requestData.Status)
	}

}
