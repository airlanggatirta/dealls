// repository/user_repository.go

package repository

import (
	"dealls/models"
	"errors"

	"gorm.io/gorm"
)

type PremiumTypeRepository struct {
	db *gorm.DB
}

func NewPremiumTypeRepository(db *gorm.DB) *PremiumTypeRepository {
	return &PremiumTypeRepository{db: db}
}

func (ptr *PremiumTypeRepository) SetDB(db *gorm.DB) {
	ptr.db = db
}

func (ptr *PremiumTypeRepository) GetPremiumTypeList() (*models.PremiumType, error) {
	if ptr.db == nil {
		return nil, errors.New("database connection not initialized")
	}

	var premiumType models.PremiumType
	if err := ptr.db.Take(&premiumType).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return &premiumType, nil
}

func (ptr *PremiumTypeRepository) GetPremiumTypeByID(premiumTypeID uint64) (*models.PremiumType, error) {
	if ptr.db == nil {
		return nil, errors.New("database connection not initialized")
	}

	var premiumType models.PremiumType
	if err := ptr.db.First(&premiumType, premiumTypeID).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return &premiumType, nil
}
