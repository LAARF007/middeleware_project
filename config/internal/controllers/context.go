package controllers

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	"net/http"

	"github.com/gofrs/uuid"
)

func ContextIDs(key string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			idStr := chi.URLParam(r, "id")
			parsedID, err := uuid.FromString(idStr)
			if err != nil {
				body, status := helpers.RespondError(&models.ErrorUnprocessableEntity{
					Message: fmt.Sprintf("cannot parse id (%s) as UUID", idStr),
				})

				w.WriteHeader(status)
				if body != nil {
					_, _ = w.Write(body)
				}
				return
			}

			ctx := context.WithValue(r.Context(), key, parsedID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
