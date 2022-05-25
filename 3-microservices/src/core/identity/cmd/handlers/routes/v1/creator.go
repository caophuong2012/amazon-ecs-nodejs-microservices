package v1

import (
	"identity/internal/controllers"
	"identity/internal/databases"

	"github.com/go-chi/chi"
)

func CreatorRoutes(r chi.Router, dbStore databases.DBStore) {
	r.Route("/creators", func(r chi.Router) {
		r.Put("/{id}", controllers.UpdateCreator(dbStore))
	})
}
