package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID                 uint64    `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Username           string    `json:"user_name" gorm:"unique;not null"`
	UserEmail          string    `json:"user_email" gorm:"unique;not null"`
	UserPassword       string    `json:"user_password" gorm:"not null"`
	UserFullname       string    `json:"user_fullname" gorm:"not null"`
	UserProfilePicture string    `json:"user_profile_picture" gorm:"sql.NullString"`
	UserVerifiedStatus uint16    `json:"user_verified_status" gorm:"default:0"`
	PremiumID          uint64    `json:"premium_id"`
	PremiumType        uint16    `json:"premium_type" gorm:"sql.NullInt16"`
	PremiumExpiry      time.Time `json:"premium_expiry" gorm:"sql.NullTime"`
	DailyStatisticID   uint64    `json:"daily_statistic_id"`
}
