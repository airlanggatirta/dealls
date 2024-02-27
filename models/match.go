package models

import (
	"gorm.io/gorm"
)

type Match struct {
	gorm.Model
	ID      uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	UserID  uint64 `json:"user_id" gorm:"not null"`
	MatchID uint64 `json:"match_id" gorm:"not null"`
}
