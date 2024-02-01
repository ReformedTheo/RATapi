package handlers

import (
	"RATapi/database"
	"RATapi/models"
	"encoding/json"
	"net/http"
	"strings"
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

	// Verifica se o código HEX começa com B (bobina) ou RA (rack)
	if strings.HasPrefix(requestData.HEX, "B") {
		// Atualiza a bobina
		database.UpdateCoilStatus(requestData.HEX, requestData.State)
	} else {
		http.Error(w, "Invalid HEX code.", http.StatusBadRequest)
	}
}
