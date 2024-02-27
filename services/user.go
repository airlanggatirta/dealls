// services/login_service.go

package services

import (
	"dealls/models"
	"dealls/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository    *repository.UserRepository
	PremiumRepository *repository.PremiumRepository
}

func NewUserService(userRepository *repository.UserRepository, premiumRepository *repository.PremiumRepository) *UserService {
	return &UserService{UserRepository: userRepository, PremiumRepository: premiumRepository}
}

func (us *UserService) Login(email, password string) error {
	user, err := us.UserRepository.FindByUserEmail(email)
	if err != nil {
		return ErrInvalidCredentials
	}

	if user.UserPassword != password {
		return ErrInvalidCredentials
	}

	return nil
}

func (us *UserService) Register(username, password, email, fullname, profile_picture string) error {

	_, err := us.UserRepository.FindByUserEmail(email)
	if err == nil {
		return ErrUsernameExists
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return err
	}

	newUser := &models.User{
		Username:           username,
		UserPassword:       hashedPassword,
		UserEmail:          email,
		UserFullname:       fullname,
		UserProfilePicture: profile_picture,
		PremiumType:        1,
		PremiumExpiry:      time.Now(),
	}

	if err := us.UserRepository.Create(newUser); err != nil {
		return err
	}

	newPremium := &models.Premium{
		PremiumType: 1,
	}

	if err := us.PremiumRepository.Create(newPremium); err != nil {
		return err
	}

	return nil
}

func (us *UserService) CheckPremiumUserStatus(user_id uint64) (*models.User, error) {
	user, err := us.UserRepository.FindUserByID(user_id)
	if err != nil {
		return nil, ErrUserNotFound
	}

	return user, nil
}

func hashPassword(password string) (string, error) {
	// Generate a bcrypt hash of the password with default cost (10)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
