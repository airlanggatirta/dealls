package models

import (
	"gorm.io/gorm"
)

type PremiumType struct {
	gorm.Model
	ID          uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	PremiumName string `json:"premium_name"`
	PremiumDesc string `json:"premium_desc" gorm:"type:text"`
}
