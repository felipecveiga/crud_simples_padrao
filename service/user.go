package service

import (
	"github.com/felipecveiga/crud_simples_padrao/model"
	"github.com/felipecveiga/crud_simples_padrao/repository"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(userReposi *repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: userReposi,
	}

}

func (us *UserService) GetUsers() ([]model.User, error) {
	users, err := us.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return users,nil
}