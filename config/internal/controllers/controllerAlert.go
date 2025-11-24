package controllers

import (
	"encoding/json"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"middleware/example/internal/services"
	"net/http"

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

func DeleteAlert(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("alertId").(uuid.UUID)

	err := services.DeleteAlert(id)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	helpers.JSON(w, http.StatusOK, map[string]string{
		"message": "Alert supprimé avec succès",
	})
}

func CreateAlert(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		AgendaId string `json:"agendaId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		body, status := helpers.RespondError(&models.ErrorBadRequest{
			Message: "Données invalides",
		})
		w.WriteHeader(status)
		_, _ = w.Write(body)
		return
	}

	if input.Email == "" {
		body, status := helpers.RespondError(&models.ErrorBadRequest{
			Message: "Champ 'email' obligatoire",
		})
		w.WriteHeader(status)
		_, _ = w.Write(body)
		return
	}

	if input.AgendaId == "" {
		body, status := helpers.RespondError(&models.ErrorBadRequest{
			Message: "Champ 'agendaId' obligatoire",
		})
		w.WriteHeader(status)
		_, _ = w.Write(body)
		return
	}

	// Valider que agendaId est un UUID valide
	agendaUUID, err := uuid.FromString(input.AgendaId)
	if err != nil {
		body, status := helpers.RespondError(&models.ErrorBadRequest{
			Message: "agendaId doit être un UUID valide",
		})
		w.WriteHeader(status)
		_, _ = w.Write(body)
		return
	}

	alert, err := services.CreateAlert(input.Email, agendaUUID)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}
	if alert == nil {
		body, status := helpers.RespondError(&models.ErrorBadRequest{
			Message: "L'agenda spécifié n'existe pas",
		})
		w.WriteHeader(status)
		_, _ = w.Write(body)
		return
	}

	helpers.JSON(w, http.StatusCreated, alert)
}

func UpdateAlert(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("alertId").(uuid.UUID)

	var input struct {
		Email    string `json:"email"`
		AgendaId string `json:"agendaId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		body, status := helpers.RespondError(&models.ErrorBadRequest{
			Message: "Données invalides",
		})
		w.WriteHeader(status)
		_, _ = w.Write(body)
		return
	}

	if input.Email == "" {
		body, status := helpers.RespondError(&models.ErrorBadRequest{
			Message: "Champ 'email' obligatoire",
		})
		w.WriteHeader(status)
		_, _ = w.Write(body)
		return
	}

	if input.AgendaId == "" {
		body, status := helpers.RespondError(&models.ErrorBadRequest{
			Message: "Champ 'agendaId' obligatoire",
		})
		w.WriteHeader(status)
		_, _ = w.Write(body)
		return
	}

	// Valider que agendaId est un UUID valide
	agendaUUID, err := uuid.FromString(input.AgendaId)
	if err != nil {
		body, status := helpers.RespondError(&models.ErrorBadRequest{
			Message: "agendaId doit être un UUID valide",
		})
		w.WriteHeader(status)
		_, _ = w.Write(body)
		return
	}

	alert, err := services.UpdateAlert(id, input.Email, agendaUUID)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	helpers.JSON(w, http.StatusOK, alert)
}
