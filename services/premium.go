// services/login_service.go

package services

import (
	"dealls/repository"
)

type PremiumService struct {
	UserRepository    *repository.UserRepository
	PremiumRepository *repository.PremiumRepository
}

func NewPremiumService(premiumRepository *repository.PremiumRepository, userRepository *repository.UserRepository) *PremiumService {
	return &PremiumService{UserRepository: userRepository, PremiumRepository: premiumRepository}
}

func (ps *PremiumService) GetPremium(premiumID uint64) error {
	_, err := ps.PremiumRepository.GetPremiumByID(premiumID)
	if err != nil {
		return ErrInvalidCredentials
	}

	return nil
}
