package v1

import (
	"creator/internal/controllers"
	"creator/internal/databases"
	"github.com/go-chi/chi"
)

func StoreFrontRoutes(r chi.Router, dbStore databases.DBStore) {
	r.Route("/store-front", func(r chi.Router) {
		r.Post("/", controllers.CreateStorefront(dbStore))
		r.Get("/{creator_id}", controllers.GetStorefront(dbStore))
	})
}
