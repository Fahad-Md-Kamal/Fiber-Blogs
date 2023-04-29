package models

import (
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
	var count int64
	query := database.DB.Model(&Users{}).Where("username = ? OR email = ?", u.Username, u.Email)
	if u.ID > 0 {
		query = query.Not("id = ?", u.ID)
	}
	err := query.Count(&count).Error
	if err != nil {
		return err.Error(), true
	}
	return "User exists with the given attribute(s)", count > 0
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

func (dbUser *Users) UpdateUser(updateDto interface{}, omitFields ...string) (*Users, error) {
	if result := database.DB.Model(dbUser).Omit(omitFields...).Updates(updateDto); result.Error != nil {
		return nil, result.Error
	}
	return dbUser, nil
}
