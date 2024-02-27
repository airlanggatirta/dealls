package models

import (
	"gorm.io/gorm"
)

type DailyStatistic struct {
	gorm.Model
	ID              uint64 `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	UserID          uint64 `json:"user_id"`
	DailySwipeCount uint64 `json:"daily_swipe_count"`
}
