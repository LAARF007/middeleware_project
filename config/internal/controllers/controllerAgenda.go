package controllers

import (
	"encoding/json"
	"middleware/example/internal/helpers"
	"middleware/example/internal/services"
	"net/http"

	"github.com/gofrs/uuid"
)

// GetAllAgendas retourne toutes les agendas
func GetAllAgendas(w http.ResponseWriter, _ *http.Request) {
	agendas, err := services.GetAllAgendas()
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(agendas)
	_, _ = w.Write(body)
}

// GetAgendaByID retourne une agenda par son ID depuis le context
func GetAgendaByID(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("agendaId").(uuid.UUID)

	agenda, err := services.GetAgendaByID(id)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(agenda)
	_, _ = w.Write(body)
}

func DeleteAgenda(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("agendaId").(uuid.UUID)

	err := services.DeleteAgenda(id)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	helpers.JSON(w, http.StatusOK, map[string]string{
		"message": "Agenda supprimé avec succès",
	})
}
