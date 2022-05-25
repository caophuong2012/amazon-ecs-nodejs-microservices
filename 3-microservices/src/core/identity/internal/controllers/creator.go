package controllers

import (
	"context"
	"encoding/json"
	"identity/internal/databases"
	"identity/internal/httpbody/request"
	"identity/internal/services"
	"identity/utils/response"
	"net/http"

	"github.com/go-chi/chi"
)

// UpdateCreator update creator
func UpdateCreator(dbStore databases.DBStore) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		creatorId := chi.URLParam(r, "id")

		var updateCreator request.UpdateCreator
		if r.Body == nil {
			response.ErrorJSON(context.Background(), w, http.StatusBadRequest, "Please send a request body", "Please send a request body")
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&updateCreator); err != nil {
			response.ErrorJSON(context.Background(), w, http.StatusBadRequest, err.Error(), err.Error())
			return
		}

		if err := services.NewCreator(dbStore).Update(creatorId, updateCreator); err != nil {
			response.ErrorJSON(context.Background(), w, http.StatusInternalServerError, err.Error(), err.Error())
		}

		response.RespondJSON(w, http.StatusOK, nil, nil)
	})
}
