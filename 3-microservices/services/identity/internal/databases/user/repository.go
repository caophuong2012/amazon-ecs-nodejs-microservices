package user

import (
	"database/sql"
	"time"

	"context"

	"identity/internal/databases/orm"
	"identity/internal/models"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Repository interface {
	Create(ctx context.Context, user models.User) (models.User, error)
}

// Management is the implementation of user
type Management struct {
	dbconn *sql.DB
}

// NewManagement initializes user
func NewManagement(dbconn *sql.DB) Repository {
	return &Management{
		dbconn: dbconn,
	}
}

func translateOrmUser(user *models.User) orm.User {
	ormUser := orm.User{
		ID:        user.ID,
		Auth0ID:   null.StringFrom(user.Auth0ID),
		CreatedAt: user.CreatedAt,
		UpdatedAt: null.TimeFrom(user.UpdatedAt),
		DeletedAt: null.TimeFrom(user.DeletedAt),
	}
	return ormUser
}

func translateUser(ormUser *orm.User) models.User {
	user := models.User{
		ID:        ormUser.ID,
		Auth0ID:   ormUser.Auth0ID.String,
		CreatedAt: ormUser.CreatedAt,
		UpdatedAt: *ormUser.UpdatedAt.Ptr(),
		DeletedAt: *ormUser.DeletedAt.Ptr(),
	}
	return user
}

// CreateUser create user
func (management *Management) Create(ctx context.Context, user models.User) (models.User, error) {

	user.CreatedAt = time.Now().UTC()

	ormUser := translateOrmUser(&user)
	if err := ormUser.Insert(ctx, management.dbconn, boil.Infer()); err != nil {
		return models.User{}, err
	}
	return translateUser(&ormUser), nil
}
