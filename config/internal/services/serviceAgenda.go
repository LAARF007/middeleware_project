package services

import (
	"database/sql"
	"fmt"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func GetAllAgendas() ([]models.Agenda, error) {
	agendas, err := repository.GetAllAgendas()
	if err != nil {
		logrus.Errorf("Erreur lors de la récupération des agendas : %s", err.Error())
		return nil, &models.ErrorGeneric{
			Message: "Une erreur est survenue lors de la récupération des agendas",
		}
	}

	return agendas, nil
}

func GetAgendaByID(id uuid.UUID) (*models.Agenda, error) {
	agenda, err := repository.GetAgendaByID(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.ErrorNotFound{
				Message: "Agenda non trouvé",
			}
		}
		logrus.Errorf("Erreur lors de la récupération de l'agenda %s : %s", id.String(), err.Error())
		return nil, &models.ErrorGeneric{
			Message: fmt.Sprintf("Une erreur est survenue lors de la récupération de l'agenda %s", id.String()),
		}
	}

	return agenda, nil
}

func DeleteAgenda(id uuid.UUID) error {
	err := repository.DeleteAgenda(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return &models.ErrorNotFound{
				Message: "Agenda non trouvé",
			}
		}

		logrus.Errorf("Erreur lors de la suppression de l'agenda %s : %s", id.String(), err.Error())
		return &models.ErrorGeneric{
			Message: fmt.Sprintf("Une erreur est survenue lors de la suppression de l'agenda %s", id.String()),
		}
	}

	return nil
}
