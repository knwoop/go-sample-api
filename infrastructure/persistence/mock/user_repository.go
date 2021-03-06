package mock

import (
	"time"

	"github.com/knwoop/go-sample-api/domain/model"
	"github.com/knwoop/go-sample-api/domain/repository"
)

type userRepository struct {
}

func NewUserRepository() repository.UserRepository {
	return &userRepository{}
}

func (ur *userRepository) Store(u *model.User) error {
	u.ID = 100
	return nil
}

func (ur *userRepository) Find(userID uint64) (*model.User, error) {

	u := &model.User{ID: userID}
	u.CreatedAt = time.Time{}
	return u, nil
}
