package models

import (
	"fmt"

	"github.com/fahad-md-kamal/fiber-blogs/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username    string `gorm:"unique;not null" json:"username"`
	Email       string `gorm:"unique;not null" json:"email"`
	Password    string `gorm:"not null" json:"password"`
	IsSuperuser bool   `gorm:"default=false;not null" json:"is_superuser"`
	IsActive    bool   `gorm:"default=true;not null" json:"is_active"`
}

func (u *Users) ValidateUserExists() (string, bool) {
	var user Users
	result := database.DB.Where("username = ? OR email = ?", u.Username, u.Email).First(&user)
	return fmt.Sprintf("User exists with username: %s OR email: %s", u.Username, u.Email), result.RowsAffected > 0
}

func (u *Users) Save() error {
	if u.ID == 0 {
		if result := database.DB.Create(&u); result.Error != nil {
			return result.Error
		}
	} else {
		if result := database.DB.Save(&u); result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func (u *Users) GeneratePasswordHash() (error, bool) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err, false
	}
	u.Password = string(hashedPassword)
	return nil, true
}

func GetUsersList(limit, offset int) ([]Users, int64, error) {
	var users []Users
	var totalCount int64

	if err := database.DB.Model(Users{}).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	if err := database.DB.Model(Users{}).Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, totalCount, nil
}

func GetUserById(userId uint) (*Users, error) {
	var user Users
	result := database.DB.First(&user, userId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
