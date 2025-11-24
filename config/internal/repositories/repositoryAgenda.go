package repositories

import (
	"database/sql"
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
)

func GetAllAgendas() ([]models.Agenda, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	rows, err := db.Query("SELECT id, ucaId, name FROM agendas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var agendas []models.Agenda
	for rows.Next() {
		var a models.Agenda
		var idStr string
		if err := rows.Scan(&idStr, &a.UcaID, &a.Name); err != nil {
			return nil, err
		}
		a.ID, _ = uuid.FromString(idStr)
		agendas = append(agendas, a)
	}

	return agendas, nil
}

func GetAgendaByID(id uuid.UUID) (*models.Agenda, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	defer helpers.CloseDB(db)

	row := db.QueryRow("SELECT id, ucaId, name FROM agendas WHERE id = ?", id.String())
	var a models.Agenda
	var idStr string
	if err := row.Scan(&idStr, &a.UcaID, &a.Name); err != nil {
		return nil, err
	}
	a.ID, _ = uuid.FromString(idStr)

	return &a, nil
}

func DeleteAgenda(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	result, err := db.Exec(`DELETE FROM agendas WHERE id = ?`, id.String())
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func CreateAgenda(a *models.Agenda) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	_, err = db.Exec(
		`INSERT INTO agendas (id, ucaId, name) VALUES (?, ?, ?)`,
		a.ID.String(), a.UcaID, a.Name,
	)

	if err != nil {
		return err
	}

	return nil
}

func UpdateAgenda(a *models.Agenda) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	defer helpers.CloseDB(db)

	result, err := db.Exec(
		`UPDATE agendas SET ucaId = ?, name = ? WHERE id = ?`,
		a.UcaID, a.Name, a.ID.String(),
	)

	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
