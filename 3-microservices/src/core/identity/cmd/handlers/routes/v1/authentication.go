package v1

import (
	"identity/internal/controllers"
	"identity/internal/databases"

	"github.com/go-chi/chi"
)

// AuthenticationRoutes adds routes for authentications
func AuthenticationRoutes(r chi.Router, dbStore databases.DBStore) {
	r.Route("/authentications", func(r chi.Router) {
		r.Route("/sign-up", func(r chi.Router) {
			r.Post("/creator", controllers.CreatorSignUp(dbStore))
		})
	})
}
