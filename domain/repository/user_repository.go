package repository

import (
	"github.com/knwoop/go-sample-api/domain/model"
)

type UserRepository interface {
	Store(u *model.User) error
	Find(userID uint64) (*model.User, error)
}
