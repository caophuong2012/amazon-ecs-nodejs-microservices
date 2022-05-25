package services

import (
	"context"
	"identity/internal/databases"
	"identity/internal/httpbody/request"
	"identity/internal/models"
)

type ICreator interface {
	Update(creatorId string, updateCreator request.UpdateCreator) error
}

type Creator struct {
	dbStore databases.DBStore
}

func NewCreator(dbStore databases.DBStore) ICreator {
	return &Creator{
		dbStore: dbStore,
	}
}

func (creator *Creator) Update(creatorId string, updateCreator request.UpdateCreator) error {
	modelCreator := models.Creator{
		ID:   creatorId,
		Type: updateCreator.Type,
	}
	return creator.dbStore.Creator.Update(context.Background(), modelCreator)
}
