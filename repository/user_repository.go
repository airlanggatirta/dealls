// repository/user_repository.go

package repository

import (
	"dealls/models"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) SetDB(db *gorm.DB) {
	ur.db = db
}

func (ur *UserRepository) FindByUserEmail(email string) (*models.User, error) {
	if ur.db == nil {
		return nil, errors.New("database connection not initialized")
	}

	var user models.User
	if err := ur.db.Where("user_email = ?", email).First(&user).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) Create(user *models.User) error {

	if ur.db == nil {
		return errors.New("database connection not initialized")
	}

	if err := ur.db.Create(user).Error; err != nil {
		return err
	}

	return nil

}

func (ur *UserRepository) FindUserByID(userID uint64) (*models.User, error) {
	if ur.db == nil {
		return nil, errors.New("database connection not initialized")
	}

	var user models.User
	if err := ur.db.First(&user, userID).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}
