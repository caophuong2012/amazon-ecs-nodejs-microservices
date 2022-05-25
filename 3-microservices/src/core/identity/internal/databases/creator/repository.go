package creator

import (
	"context"
	"database/sql"
	"identity/internal/databases/orm"
	"identity/internal/models"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Repository interface {
	Create(ctx context.Context, user models.Creator) (models.Creator, error)
	Update(ctx context.Context, newCreator models.Creator) error
}

// Management is the implementation of creator
type Management struct {
	dbconn *sql.DB
}

// NewManagement initializes creator
func NewManagement(dbconn *sql.DB) Repository {
	return &Management{
		dbconn: dbconn,
	}
}

func translateOrmCreator(creator *models.Creator) orm.Creator {
	ormCreator := orm.Creator{
		ID:        creator.ID,
		UserID:    null.StringFrom(creator.UserID),
		Type:      null.StringFrom(creator.Type),
		CreatedAt: creator.CreatedAt,
		UpdatedAt: null.TimeFrom(creator.UpdatedAt),
		DeletedAt: null.TimeFrom(creator.DeletedAt),
	}
	return ormCreator
}

func translateCreator(ormCreator *orm.Creator) models.Creator {
	creator := models.Creator{
		ID:        ormCreator.ID,
		UserID:    ormCreator.UserID.String,
		Type:      ormCreator.Type.String,
		CreatedAt: ormCreator.CreatedAt,
		UpdatedAt: *ormCreator.UpdatedAt.Ptr(),
		DeletedAt: *ormCreator.DeletedAt.Ptr(),
	}
	return creator
}

// Create creator
func (management *Management) Create(ctx context.Context, creator models.Creator) (models.Creator, error) {

	ormCreator := translateOrmCreator(&creator)
	if err := ormCreator.Insert(ctx, management.dbconn, boil.Infer()); err != nil {
		return models.Creator{}, err
	}
	return translateCreator(&ormCreator), nil
}

// Update type
func (management *Management) Update(ctx context.Context, newCreator models.Creator) error {
	ormCreator, err := orm.Creators(orm.CreatorWhere.ID.EQ(newCreator.ID)).One(ctx, management.dbconn)
	if err != nil {
		return err
	}

	ormCreator.Type = null.StringFrom(newCreator.Type)

	_, err = ormCreator.Update(ctx, management.dbconn, boil.Infer())
	return err
}
