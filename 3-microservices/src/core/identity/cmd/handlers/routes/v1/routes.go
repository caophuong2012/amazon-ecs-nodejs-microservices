package v1

import (
	"identity/internal/databases"
	"identity/internal/enums"
	"identity/internal/middlewares"
	"net/http"

	"github.com/go-chi/chi"
)

func Routes(r chi.Router, dbStore databases.DBStore) {
	r.Route("/v1", func(r chi.Router) {
		r.Use(middlewares.EnsureValidToken())

		// this is for test HasScope middleware
		r.With(middlewares.HasScope(enums.ReadProfile)).
			Get("/profile", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("OK"))
			})

		AuthenticationRoutes(r, dbStore)
		CreatorRoutes(r, dbStore)
	})
}
