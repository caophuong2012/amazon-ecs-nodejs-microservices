package services

import (
	"context"
	auth0Service "identity/external/auth0"
	"identity/internal/databases"
	"identity/internal/models"
	"identity/utils"

	"github.com/gofrs/uuid"
	"gopkg.in/auth0.v5"
	"gopkg.in/auth0.v5/management"
)

type IAuthService interface {
	CreatorSignUp(Auth0ID string) (models.User, error)
}

type AuthService struct {
	dbStore databases.DBStore
}

func NewAuthService(dbStore databases.DBStore) IAuthService {
	return &AuthService{
		dbStore: dbStore,
	}
}

// CreatorSignUp will create an user and creator
func (authService *AuthService) CreatorSignUp(Auth0ID string) (models.User, error) {
	userId, _ := uuid.NewV4()
	user := models.User{ID: userId.String(), Auth0ID: Auth0ID}
	user, err := authService.dbStore.User.Create(context.Background(), user)

	if err != nil {
		return models.User{}, err
	}

	creatorId, _ := uuid.NewV4()
	creator := models.Creator{ID: creatorId.String(), UserID: user.ID}
	creator, err = authService.dbStore.Creator.Create(context.Background(), creator)

	if err != nil {
		return models.User{}, err
	}

	creatorRoleId := utils.GetWithDefault("AUTH0_CREATOR_ROLE_ID", "")
	role := &management.Role{
		ID: auth0.String(creatorRoleId),
	}

	err = auth0Service.AssignRolesToUser(user.Auth0ID, role)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
