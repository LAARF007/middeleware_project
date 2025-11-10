package controllers

import (
	"encoding/json"
	"net/http"
	"projetgoo/internal/helpers"
	"projetgoo/internal/services"

	"github.com/gofrs/uuid"
)

func GetAllAlerts(w http.ResponseWriter, _ *http.Request) {
	alerts, err := services.GetAllAlerts()
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(alerts)
	_, _ = w.Write(body)
}

func GetAlertByID(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("alertId").(uuid.UUID)

	alert, err := services.GetAlertByID(id)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(alert)
	_, _ = w.Write(body)
}
