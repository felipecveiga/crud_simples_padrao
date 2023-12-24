package repository

import (


	"github.com/felipecveiga/crud_simples_padrao/model"
	"gorm.io/gorm"
)


type UserRepository struct {
	DB gorm.DB
}

func NewUserRepository (db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: *db,
	}
}

func (r UserRepository) FindAll() ([]model.User, error) {
	var users []model.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetUserByIDFromDB(userID string) (*model.User, error) {
	var user model.User
	if err := r.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}


func (r UserRepository) UpdateUserInDB(userID string, userData model.User) (*model.User, error) {
	//CRIAR UM MAP COM OS CAMPOS QUE DEVEM SER ATUALIZADOS
	updateFields := map[string]interface{} {
		"name": userData.Name,
		"idade": userData.Idade,
		"email": userData.Email,
	}

	//REALIZAR A ATUALIZACAO APENAS DOS CAMPOS ESPECIFICADO NO BANCO DE DADOS USANDO O GORM
	var user model.User
	if err := r.DB.Model(&user).Where("id = ?", userID).Updates(updateFields).Error; err != nil {
		return nil, err
	}
	return &user, nil
}