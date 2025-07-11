package userservice

import (
	"sync"
	"sync/atomic"

	"github.com/AbhinitKumarRai/user-service/pkg/model"
)

type UserService struct {
	users         sync.Map
	userIDCounter int64
}

func NewUserService() *UserService {
	return &UserService{
		users:         sync.Map{},
		userIDCounter: 0,
	}
}

func (r *UserService) Create(user model.User) model.User {
	user.ID = int(atomic.AddInt64(&(r.userIDCounter), 1))

	r.users.Store(user.ID, &user)
	return user
}

func (r *UserService) GetByID(id int) (model.User, bool) {
	val, ok := r.users.Load(id)
	if !ok {
		return model.User{}, false
	}
	return *(val.(*model.User)), true
}

func (r *UserService) Update(id int, update model.User) (model.User, bool) {
	val, exists := r.users.Load(id)
	if !exists {
		return model.User{}, false
	}

	// Type assertion from `any` to `*model.User`
	user, ok := val.(*model.User)
	if !ok {
		return model.User{}, false // type mismatch — should not happen in normal usage
	}

	if user.Name != update.Name {
		user.Name = update.Name
	}

	if user.PasswordHash != update.PasswordHash {
		user.PasswordHash = update.PasswordHash
	}

	return *user, true
}

func (r *UserService) Delete(id int) bool {
	_, exists := r.users.Load(id)
	if exists {
		r.users.Delete(id)
		return true
	}
	return false
}

func (r *UserService) FindByEmail(email string) (model.User, bool) {
	var result model.User
	found := false
	r.users.Range(func(_, value any) bool {
		user := value.(*model.User)
		if user.Email == email {
			result = *user
			found = true
			return false
		}
		return true
	})
	return result, found
}
