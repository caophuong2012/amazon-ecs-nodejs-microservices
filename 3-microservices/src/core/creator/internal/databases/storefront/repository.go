package storefront

import (
	"context"
	"creator/internal/databases/orm"
	"creator/internal/models"
	"database/sql"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
)

type Repository interface {
	Create(ctx context.Context, storeFront models.StoreFront) (models.StoreFront, error)
	Retrieve(ctx context.Context, creatorID int64) (models.StoreFront, error)
}

// Management is the implementation of storefront
type Management struct {
	dbConn *sql.DB
}

// NewManagement initializes storefront
func NewManagement(dbConn *sql.DB) Repository {
	return &Management{
		dbConn: dbConn,
	}
}

// Retrieve get store front by creator id
func (management *Management) Retrieve(ctx context.Context, creatorID int64) (models.StoreFront, error) {
	storeFront, err := orm.StoreFronts(
		orm.StoreFrontWhere.CreatorID.EQ(null.Int64From(creatorID)),
	).One(ctx, management.dbConn)
	if err != nil {
		return models.StoreFront{}, err
	}

	return translateStoreFront(storeFront), nil
}

// Create storefront
func (management *Management) Create(ctx context.Context, storeFront models.StoreFront) (models.StoreFront, error) {
	ormStoreFront := translateOrmStoreFront(&storeFront)
	if err := ormStoreFront.Insert(ctx, management.dbConn, boil.Infer()); err != nil {
		return models.StoreFront{}, err
	}

	return translateStoreFront(&ormStoreFront), nil
}

func translateOrmStoreFront(storeFront *models.StoreFront) orm.StoreFront {
	ormStoreFront := orm.StoreFront{
		CreatorID:                     null.Int64From(storeFront.CreatorID),
		StoreName:                     storeFront.StoreName,
		StoreURL:                      storeFront.StoreUrl,
		StoreLogoFilename:             storeFront.StoreLogoFileName,
		StorefrontBannerVideoFilename: null.StringFrom(storeFront.StoreFrontBannerVideoFileName),
		StorefrontTitleHeader:         null.StringFrom(storeFront.StoreFrontTitleHeader),
		StorefrontSubHeader:           null.StringFrom(storeFront.StoreFrontSubHeader),
		PromotionalMessage:            null.StringFrom(storeFront.PromotionalMessage),
	}
	return ormStoreFront
}

func translateStoreFront(ormStoreFront *orm.StoreFront) models.StoreFront {
	storeFront := models.StoreFront{
		CreatorID:                     ormStoreFront.CreatorID.Int64,
		StoreName:                     ormStoreFront.StoreName,
		StoreUrl:                      ormStoreFront.StoreURL,
		StoreLogoFileName:             ormStoreFront.StoreLogoFilename,
		StoreFrontBannerVideoFileName: ormStoreFront.StorefrontBannerVideoFilename.String,
		StoreFrontTitleHeader:         ormStoreFront.StorefrontTitleHeader.String,
		StoreFrontSubHeader:           ormStoreFront.StorefrontSubHeader.String,
		PromotionalMessage:            ormStoreFront.PromotionalMessage.String,
	}
	return storeFront
}
