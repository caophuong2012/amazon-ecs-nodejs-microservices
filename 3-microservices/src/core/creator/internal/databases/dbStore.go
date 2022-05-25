package databases

import "creator/internal/databases/storefront"

type DBStore struct {
	Storefront storefront.Repository
}
