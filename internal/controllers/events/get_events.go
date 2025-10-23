package events

import (
	"encoding/json"
	"middleware/example/internal/helpers"
	"middleware/example/internal/repositories/events"
	"net/http"
)

func GetEvents(w http.ResponseWriter, _ *http.Request) {
	// calling service
	events, err := events.GetAllEvents()
	if err != nil {
		body, status := helpers.RespondError(err)
		w.WriteHeader(status)
		if body != nil {
			_, _ = w.Write(body)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(events)
	_, _ = w.Write(body)
	return
}
