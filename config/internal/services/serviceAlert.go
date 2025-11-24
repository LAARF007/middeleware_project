package services

import (
	"database/sql"
	"fmt"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func GetAllAlerts() ([]models.Alert, error) {
	alerts, err := repository.GetAllAlerts()
	if err != nil {
		logrus.Errorf("Erreur lors de la récupération des alerts : %s", err.Error())
		return nil, &models.ErrorGeneric{
			Message: "Une erreur est survenue lors de la récupération des alerts",
		}
	}

	return alerts, nil
}

func GetAlertByID(id uuid.UUID) (*models.Alert, error) {
	alert, err := repository.GetAlertByID(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.ErrorNotFound{
				Message: "Alert non trouvée",
			}
		}
		logrus.Errorf("Erreur lors de la récupération de l'alert %s : %s", id.String(), err.Error())
		return nil, &models.ErrorGeneric{
			Message: fmt.Sprintf("Une erreur est survenue lors de la récupération de l'alert %s", id.String()),
		}
	}

	return alert, nil
}

func DeleteAlert(id uuid.UUID) error {
	err := repository.DeleteAlert(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return &models.ErrorNotFound{
				Message: "Alert non trouvé",
			}
		}

		logrus.Errorf("Erreur lors de la suppression de l'alert %s : %s", id.String(), err.Error())
		return &models.ErrorGeneric{
			Message: fmt.Sprintf("Une erreur est survenue lors de la suppression de l'alert %s", id.String()),
		}
	}

	return nil
}

func CreateAlert(email string, agendaId uuid.UUID) (*models.Alert, error) {

	// --- Vérifier que l'agenda existe ---
	exists, err := repository.GetAgendaByID(agendaId)
	if err != nil {
		logrus.Errorf("Erreur vérification agenda %s : %s", agendaId.String(), err.Error())
		return nil, &models.ErrorGeneric{
			Message: "Erreur lors de la vérification de l'agenda",
		}
	}

	if exists == nil {
		return nil, &models.ErrorGeneric{
			Message: "L'agenda spécifié n'existe pas",
		}
	}

	// generate uuid
	newID, err := uuid.NewV4()
	if err != nil {
		return nil, &models.ErrorGeneric{
			Message: "Impossible de générer un UUID",
		}
	}

	alert := &models.Alert{
		ID:       newID, // ← UUID auto généré
		Email:    email,
		AgendaID: agendaId,
	}

	err = repository.CreateAlert(alert)
	if err != nil {

		logrus.Errorf("Erreur création alert %s : %s", newID.String(), err.Error())
		return nil, &models.ErrorGeneric{
			Message: "Erreur lors de la création de l'alert",
		}
	}

	return alert, nil
}

func UpdateAlert(id uuid.UUID, email string, agendaId uuid.UUID) (*models.Alert, error) {
	// Vérifier que l'alert existe
	_, err := repository.GetAlertByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &models.ErrorNotFound{
				Message: "Alert non trouvé",
			}
		}
		logrus.Errorf("Erreur lors de la récupération de l'alert %s : %s", id.String(), err.Error())
		return nil, &models.ErrorGeneric{
			Message: "Erreur lors de la vérification de l'alert",
		}
	}

	alert := &models.Alert{
		ID:       id,
		Email:    email,
		AgendaID: agendaId,
	}

	err = repository.UpdateAlert(alert)
	if err != nil {
		logrus.Errorf("Erreur mise à jour alert %s : %s", id.String(), err.Error())
		return nil, &models.ErrorGeneric{
			Message: "Erreur lors de la mise à jour de l'alert",
		}
	}

	return alert, nil
}
