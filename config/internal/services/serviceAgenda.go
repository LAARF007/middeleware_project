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

func CreateAgenda(ucaId int, name string) (*models.Agenda, error) {
	newID, err := uuid.NewV4()
	if err != nil {
		return nil, &models.ErrorGeneric{
			Message: "Impossible de générer un UUID",
		}
	}

	agenda := &models.Agenda{
		ID:    newID, // ← UUID auto généré
		UcaID: ucaId,
		Name:  name,
	}

	err = repository.CreateAgenda(agenda)
	if err != nil {

		logrus.Errorf("Erreur création agenda %s : %s", newID.String(), err.Error())
		return nil, &models.ErrorGeneric{
			Message: "Erreur lors de la création de l'agenda",
		}
	}

	return agenda, nil
}

func UpdateAgenda(id uuid.UUID, ucaId int, name string) (*models.Agenda, error) {
	// Vérifier que l'agenda existe
	_, err := repository.GetAgendaByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &models.ErrorNotFound{
				Message: "Agenda non trouvé",
			}
		}
		logrus.Errorf("Erreur lors de la récupération de l'agenda %s : %s", id.String(), err.Error())
		return nil, &models.ErrorGeneric{
			Message: "Erreur lors de la vérification de l'agenda",
		}
	}

	agenda := &models.Agenda{
		ID:    id,
		UcaID: ucaId,
		Name:  name,
	}

	err = repository.UpdateAgenda(agenda)
	if err != nil {
		logrus.Errorf("Erreur mise à jour agenda %s : %s", id.String(), err.Error())
		return nil, &models.ErrorGeneric{
			Message: "Erreur lors de la mise à jour de l'agenda",
		}
	}

	return agenda, nil
}
