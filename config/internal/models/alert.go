package models

import "github.com/gofrs/uuid"

type Alert struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	AgendaID uuid.UUID `json:"agendaId"`
}
