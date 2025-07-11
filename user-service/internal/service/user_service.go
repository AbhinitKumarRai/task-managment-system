package service

import (
	manager "github.com/AbhinitKumarRai/user-service/internal/usermanager"
	"github.com/AbhinitKumarRai/user-service/pkg/model"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	manager *manager.UserManager
}

func NewUserService(manager *manager.UserManager) *UserService {
	return &UserService{manager: manager}
}

func (s *UserService) Register(name, email, password string) (model.User, error) {
	user := model.User{Name: name, Email: email, PasswordHash: password}
	return s.manager.Create(user)
}

func (s *UserService) Login(email, password string) (model.User, error) {
	user, err := s.manager.FindByEmail(email)
	if err != nil {
		return model.User{}, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (s *UserService) GetByID(id int) (model.User, error) {
	return s.manager.GetByID(id)
}

func (s *UserService) Update(id int, update model.User) (model.User, error) {
	return s.manager.Update(id, update)
}

func (s *UserService) Delete(id int) error {
	return s.manager.Delete(id)
}
