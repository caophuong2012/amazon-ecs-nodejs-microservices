package services

import (
	"context"
	"creator/internal/databases"
	"creator/internal/httpbody/request"
	"creator/internal/models"
)

type IStoreFront interface {
	Create(createStoreFront request.CreateStoreFront) (models.StoreFront, error)
	Retrieve(creatorID int64) (models.StoreFront, error)
}

type StoreFront struct {
	dbStore databases.DBStore
}

func NewStoreFront(dbStore databases.DBStore) IStoreFront {
	return &StoreFront{
		dbStore: dbStore,
	}
}

func (storeFront *StoreFront) Retrieve(creatorID int64) (models.StoreFront, error) {
	storeFrontResult, err := storeFront.dbStore.Storefront.Retrieve(context.Background(), creatorID)
	if err != nil {
		return models.StoreFront{}, err
	}

	return storeFrontResult, err
}

func (storeFront *StoreFront) Create(createStoreFront request.CreateStoreFront) (models.StoreFront, error) {
	modelStoreFront := models.StoreFront{
		CreatorID:             createStoreFront.CreatorID,
		StoreName:             createStoreFront.StoreName,
		StoreUrl:              createStoreFront.StoreUrl,
		StoreFrontTitleHeader: createStoreFront.StoreFrontTitleHeader,
		StoreFrontSubHeader:   createStoreFront.StoreFrontSubHeader,
		PromotionalMessage:    createStoreFront.PromotionalMessage,
		StoreLogoFileName:     createStoreFront.StoreLogoFileName,
	}

	if createStoreFront.IsUrl == "true" {
		modelStoreFront.StoreFrontBannerVideoFileName = createStoreFront.StoreFrontBannerVideoUrl
	}

	return storeFront.dbStore.Storefront.Create(context.Background(), modelStoreFront)
}
