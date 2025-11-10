package services

import (
	"database/sql"
	"fmt"
	"projetgoo/internal/models"
	repository "projetgoo/internal/repositories"

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
