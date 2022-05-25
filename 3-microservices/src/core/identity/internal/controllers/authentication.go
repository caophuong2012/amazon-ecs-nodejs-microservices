package controllers

import (
	"encoding/json"
	"identity/internal/databases"
	"identity/internal/services"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

// CreatorSignUp will create an user and creator
func CreatorSignUp(dbStore databases.DBStore) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
		auth0Id := token.RegisteredClaims.Subject

		user, err := services.NewAuthService(dbStore).CreatorSignUp(auth0Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		payload, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	})
}
