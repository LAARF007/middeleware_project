package models

import "github.com/gofrs/uuid"

type Agenda struct {
	ID    uuid.UUID `json:"id"`
	UcaID int       `json:"ucaId"`
	Name  string    `json:"name"`
}
