package database

import (
	"dealls/config"
	"dealls/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() *gorm.DB {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %s", err)
	}

	dsn := cfg.Database.Username + ":" + cfg.Database.Password + "@tcp(" + cfg.Database.Host + ":" + cfg.Database.Port + ")/" + cfg.Database.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	db.AutoMigrate(&models.User{}, &models.Premium{}, &models.PremiumType{}, &models.Payment{}, &models.DailyStatistic{}, &models.Match{})

	return db
}
