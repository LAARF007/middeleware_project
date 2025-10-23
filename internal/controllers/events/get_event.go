package events

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/repositories/events"
	"net/http"
)

func GetEvent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	eventId, _ := ctx.Value("eventId").(uuid.UUID) // getting key set in context.go

	event, err := events.GetEventById(eventId)
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(event)
	_, _ = w.Write(body)
	return
}
