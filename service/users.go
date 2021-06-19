package service

import (
	"github.com/mkaiho/go-deploy-sample/domain/model/user"
	"github.com/mkaiho/go-deploy-sample/repository"
)

type UserService interface {
	Find(id string) (*user.User, error)
	Create(user user.User) error
}

type userService struct {
	userRepository *repository.UserRepository
}

func (service *userService) Find(id string) (*user.User, error) {
	return (*service.userRepository).Find(id)
}

func (service *userService) Create(user user.User) error {
	return (*service.userRepository).Create(&user)
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: &userRepository,
	}
}
