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

func (us *UserService) GetAllUsers() ([]model.User, error) {
	users, err := us.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) GetUserByID(userID string) (*model.User, error) {
	user, err := s.UserRepository.GetUserByIDFromDB(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) UserUpdate(userID string, userData model.User) (*model.User, error) {
	//VERIFICAR SE O USUARIO EXISTE
	_, err := s.UserRepository.GetUserByIDFromDB(userID)
	if err != nil {
		return nil, err
	}

	//ATUALIZAR OS DADOS DO USUARIO COM OS NOVOS DADOS RECEBIDOS
	userUpdate, err := s.UserRepository.UpdateUserInDB(userID, userData)
	if err != nil {
		return nil, err
	}

	return userUpdate, nil

}
