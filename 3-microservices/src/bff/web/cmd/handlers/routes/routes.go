package routes

import (
	"net/http"
	"web/internal/controllers"

	"github.com/go-chi/chi"
)

// Route function
func Route() http.Handler {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Get("/health-check", controllers.HealthCheck)
	})

	return r
}
