package models

import (
	"gorm.io/gorm"
)

type Premium struct {
	gorm.Model
	ID          uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	PremiumType uint16 `json:"premium_type"`
}
