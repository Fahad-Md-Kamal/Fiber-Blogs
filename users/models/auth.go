package models

import (
	"log"

	"github.com/fahad-md-kamal/fiber-blogs/database"
	"gorm.io/gorm"
)

type BlacklistedTokens struct {
	gorm.Model
	Token string `gorm:"uniqueIndex" json:"token"`
}

func (t *BlacklistedTokens) IsTokenBlacklisted() bool {
	var count int64
	if err := database.DB.Where("token = ?", t.Token).Find(&t).Count(&count).Error; err == nil {
		return count > 0
	}
	return count > 0
}

func (t *BlacklistedTokens) BlacklistToken() bool {
	if err := database.DB.Model(&t).Create(&t).Error; err != nil {
		log.Printf("Error making token as blacklisted %s", err.Error())
		return false
	}
	return true
}
