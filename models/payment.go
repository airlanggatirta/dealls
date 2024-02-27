package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	ID          uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	PremiumType uint16 `json:"premium_type"`
}
