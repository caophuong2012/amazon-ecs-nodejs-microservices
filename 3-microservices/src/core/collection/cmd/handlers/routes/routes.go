package routes

import (
	"collection/internal/controllers"
	"collection/internal/databases"
	"net/http"

	"github.com/go-chi/chi"
)

// Route function
func Route(dbStore databases.DBStore) http.Handler {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Get("/health-check", controllers.HealthCheck)
	})

	return r
}
