// repository/user_repository.go

package repository

import (
	"dealls/models"
	"errors"

	"gorm.io/gorm"
)

type PremiumRepository struct {
	db *gorm.DB
}

func NewPremiumRepository(db *gorm.DB) *PremiumRepository {
	return &PremiumRepository{db: db}
}

func (pr *PremiumRepository) SetDB(db *gorm.DB) {
	pr.db = db
}

func (pr *PremiumRepository) Create(premium *models.Premium) error {

	if pr.db == nil {
		return errors.New("database connection not initialized")
	}

	if err := pr.db.Table("premium").Create(premium).Error; err != nil {
		return err
	}

	return nil

}

func (pr *PremiumRepository) GetPremiumList() (*models.Premium, error) {
	if pr.db == nil {
		return nil, errors.New("database connection not initialized")
	}

	var premium models.Premium
	if err := pr.db.Take(&premium).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return &premium, nil
}

func (pr *PremiumRepository) GetPremiumByID(premiumID uint64) (*models.Premium, error) {
	if pr.db == nil {
		return nil, errors.New("database connection not initialized")
	}

	var premium models.Premium
	if err := pr.db.First(&premium, premiumID).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return &premium, nil
}
