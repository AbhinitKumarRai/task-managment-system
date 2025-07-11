package usermanager

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/AbhinitKumarRai/user-service/pkg/model"
	"golang.org/x/crypto/bcrypt"
)

type UserManager struct {
	users         sync.Map
	userIDCounter int64
}

func NewUserManager() *UserManager {
	return &UserManager{
		users:         sync.Map{},
		userIDCounter: 0,
	}
}

func (r *UserManager) Create(user model.User) (model.User, error) {
	if user.Email == "" || user.Name == "" || user.PasswordHash == "" {
		return model.User{}, fmt.Errorf("user data must contain name, email and password")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, err
	}

	user.ID = int(atomic.AddInt64(&(r.userIDCounter), 1))

	user.PasswordHash = string(hash)
	r.users.Store(user.ID, &user)
	return user, nil
}

func (r *UserManager) GetByID(id int) (model.User, error) {
	val, ok := r.users.Load(id)
	if !ok {
		return model.User{}, fmt.Errorf("no user with id: %d found", id)
	}
	return *(val.(*model.User)), nil
}

func (r *UserManager) Update(id int, update model.User) (model.User, error) {
	val, exists := r.users.Load(id)
	if !exists {
		return model.User{}, fmt.Errorf("no user with id: %d found", id)
	}

	// Type assertion from `any` to `*model.User`
	user, ok := val.(*model.User)
	if !ok {
		return model.User{}, fmt.Errorf("unable to get user data") // type mismatch — should not happen in normal usage
	}

	if user.Name != update.Name {
		user.Name = update.Name
	}

	if user.PasswordHash != update.PasswordHash {
		user.PasswordHash = update.PasswordHash
	}

	return *user, nil
}

func (r *UserManager) Delete(id int) error {
	_, exists := r.users.Load(id)
	if exists {
		r.users.Delete(id)
	}
	return nil
}

func (r *UserManager) FindByEmail(email string) (model.User, error) {
	var result model.User
	r.users.Range(func(_, value any) bool {
		user := value.(*model.User)
		if user.Email == email {
			result = *user
			return false
		}
		return true
	})
	return result, nil
}
