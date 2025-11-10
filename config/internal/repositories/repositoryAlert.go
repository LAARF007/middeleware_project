package repositories

import (
	"projetgoo/internal/helpers"
	"projetgoo/internal/models"

	"github.com/gofrs/uuid"
)

func GetAllAlerts() ([]models.Alert, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	rows, err := db.Query("SELECT id, email, agendaId FROM alerts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alerts []models.Alert
	for rows.Next() {
		var al models.Alert
		var idStr, agendaIDStr string
		if err := rows.Scan(&idStr, &al.Email, &agendaIDStr); err != nil {
			return nil, err
		}
		al.ID, _ = uuid.FromString(idStr)
		al.AgendaID, _ = uuid.FromString(agendaIDStr)
		alerts = append(alerts, al)
	}

	return alerts, nil
}

func GetAlertByID(id uuid.UUID) (*models.Alert, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	row := db.QueryRow("SELECT id, email, agendaId FROM alerts WHERE id = ?", id.String())
	var al models.Alert
	var idStr, agendaIDStr string
	if err := row.Scan(&idStr, &al.Email, &agendaIDStr); err != nil {
		return nil, err
	}
	al.ID, _ = uuid.FromString(idStr)
	al.AgendaID, _ = uuid.FromString(agendaIDStr)

	return &al, nil
}
