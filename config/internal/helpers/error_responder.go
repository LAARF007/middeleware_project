package helpers

import (
	"encoding/json"
	"middleware/example/internal/models"
	"net/http"
)

// RespondError convertit une erreur custom en code HTTP et body JSON
func RespondError(err error) ([]byte, int) {
	switch e := err.(type) {
	case *models.ErrorNotFound:
		body, _ := json.Marshal(e)
		return body, http.StatusNotFound
	case *models.ErrorUnprocessableEntity:
		body, _ := json.Marshal(e)
		return body, http.StatusUnprocessableEntity
	case *models.ErrorGeneric:
		body, _ := json.Marshal(e)
		return body, http.StatusInternalServerError
	default:
		// erreur inconnue
		body, _ := json.Marshal(&models.ErrorGeneric{Message: e.Error()})
		return body, http.StatusInternalServerError
	}
}
