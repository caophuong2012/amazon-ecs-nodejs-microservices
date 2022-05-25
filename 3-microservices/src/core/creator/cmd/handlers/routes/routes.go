package routes

import (
	v1 "creator/cmd/handlers/routes/v1"
	"creator/internal/controllers"
	"creator/internal/databases"
	"net/http"

	"github.com/go-chi/chi"
)

// Route function
func Route(dbStore databases.DBStore) http.Handler {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Get("/health-check", controllers.HealthCheck)
		v1.StoreFrontRoutes(r, dbStore)
	})

	return r
}
