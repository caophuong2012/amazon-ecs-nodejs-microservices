package databases

import (
	"identity/internal/databases/creator"
	"identity/internal/databases/user"
)

type DBStore struct {
	User    user.Repository
	Creator creator.Repository
}
