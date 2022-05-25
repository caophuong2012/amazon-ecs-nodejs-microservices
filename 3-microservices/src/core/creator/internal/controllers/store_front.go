package controllers

import (
	"context"
	"creator/internal/databases"
	"creator/internal/httpbody/request"
	"creator/internal/models"
	"creator/internal/services"
	"creator/utils/response"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type CreateStorefrontResponse struct {
	Status string `json:"status"`
}

func CreateStorefront(dbStore databases.DBStore) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var createStoreFront request.CreateStoreFront
		creatorID := r.FormValue("creator_id")
		if creatorID == "" {
			response.ErrorJSON(context.Background(), w, http.StatusBadRequest, "creator_id_empty", "creator id cannot be empty")
			return
		}
		creatorIDParsedToInt, err := strconv.ParseInt(creatorID, 10, 64)
		if err != nil {
			response.ErrorJSON(context.Background(), w, http.StatusBadRequest, "creator_id_invalid", "creator id invalid")
			return
		}

		storeName := r.FormValue("store_name")
		if storeName == "" {
			response.ErrorJSON(context.Background(), w, http.StatusBadRequest, "store_name_empty", "store name cannot be empty")
			return
		}

		storeUrl := r.FormValue("store_url")
		if storeUrl == "" {
			response.ErrorJSON(context.Background(), w, http.StatusBadRequest, "store_url_empty", "store url cannot be empty")
			return
		}

		storeLogoFile, storeLogoHeader, err := r.FormFile("store_logo_file")
		if storeLogoHeader == nil {
			response.ErrorJSON(context.Background(), w, http.StatusBadRequest, "header_missing", "header is missing")
			return
		}
		if err != nil {
			response.ErrorJSON(context.Background(), w, http.StatusBadRequest, "store_logo_empty", "store logo cannot be empty")
			return
		}

		isUrl := r.FormValue("is_url")
		if isUrl == "" {
			response.ErrorJSON(context.Background(), w, http.StatusBadRequest, "is_url_empty", "is url cannot be empty")
			return
		}

		storeFrontTitleHeader := r.FormValue("storefront_title_header")
		storeFrontSubHeader := r.FormValue("storefront_sub_header")
		promotionalMessage := r.FormValue("promotional_message")

		createStoreFront.CreatorID = creatorIDParsedToInt
		createStoreFront.StoreName = storeName
		createStoreFront.StoreUrl = storeUrl
		createStoreFront.StoreFrontTitleHeader = storeFrontTitleHeader
		createStoreFront.StoreFrontSubHeader = storeFrontSubHeader
		createStoreFront.PromotionalMessage = promotionalMessage
		createStoreFront.StoreLogoFile = storeLogoFile
		createStoreFront.StoreLogoFileName = storeLogoHeader.Filename

		if isUrl == "true" {
			storefrontBannerVideoUrl := r.FormValue("storefront_banner_video_url")
			createStoreFront.IsUrl = isUrl
			createStoreFront.StoreFrontBannerVideoUrl = storefrontBannerVideoUrl
		} else {
			storeFrontBannerVideoFile, storeFrontBannerVideoHeader, err := r.FormFile("storefront_banner_video_file")
			if storeFrontBannerVideoHeader == nil {
				response.ErrorJSON(context.Background(), w, http.StatusBadRequest, "header_missing", "header is missing")
				return
			}
			if err != nil {
				response.ErrorJSON(context.Background(), w, http.StatusBadRequest, "store_logo_empty", "store logo cannot be empty")
				return
			}
			createStoreFront.IsUrl = isUrl
			createStoreFront.StoreFrontBannerVideoFile = storeFrontBannerVideoFile
			createStoreFront.StoreFrontBannerVideoFileName = storeFrontBannerVideoHeader.Filename
		}

		_, err = services.NewStoreFront(dbStore).Create(createStoreFront)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		resp := CreateStorefrontResponse{
			Status: "success",
		}
		response.RespondJSON(w, http.StatusOK, resp, nil)
	})
}

func GetStorefront(dbStore databases.DBStore) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "creator_id")
		if id == "" {
			response.ErrorJSON(context.Background(), w, http.StatusBadRequest, "creator_id_empty", "creator id cannot be empty")
			return
		}

		creatorID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.ErrorJSON(context.Background(), w, http.StatusBadRequest, "creator_id_invalid", "creator id invalid")
			return
		}

		result, err := services.NewStoreFront(dbStore).Retrieve(creatorID)
		if err != nil {
			response.ErrorJSON(context.Background(), w, http.StatusInternalServerError, "unable_to_retrieve_store_front", err.Error())
			return
		}

		response.RespondJSON(w, http.StatusOK, translateToStoreFrontResponse(result), nil)
	})
}

func translateToStoreFrontResponse(storeFront models.StoreFront) StoreFrontResponse {
	storeFrontResponse := StoreFrontResponse{
		CreatorID:                     storeFront.CreatorID,
		StoreName:                     storeFront.StoreName,
		StoreUrl:                      storeFront.StoreUrl,
		StoreFrontTitleHeader:         storeFront.StoreFrontTitleHeader,
		StoreFrontSubHeader:           storeFront.StoreFrontSubHeader,
		PromotionalMessage:            storeFront.PromotionalMessage,
		StoreLogoFileName:             storeFront.StoreLogoFileName,
		StoreFrontBannerVideoFileName: storeFront.StoreFrontBannerVideoFileName,
	}

	return storeFrontResponse
}
