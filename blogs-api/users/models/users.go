package models

import (
	"log"

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

type UserCheckParams struct {
	UserId   uint
	Username string
	Email    string
}

func GetUsersList(limit, offset int) ([]Users, int64, error) {
	var users []Users
	var totalCount int64

	if err := database.DB.Model(Users{}).Count(&totalCount).Error; err != nil {
		log.Printf("Failed to get users list count: %s", err.Error())
		return nil, 0, err
	}

	if err := database.DB.Model(Users{}).Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		log.Printf("Failed to get users list with Limit: %d | Offset: %d | Error: %s", limit, offset, err.Error())
		return nil, 0, err
	}

	return users, totalCount, nil
}

func GetUserById(userId uint) (*Users, error) {
	var user Users
	result := database.DB.First(&user, userId)
	if result.Error != nil {
		log.Printf("Failed to get user by ID: %d | Error: %s", userId, result.Error.Error())
		return nil, result.Error
	}
	return &user, nil
}

func (u *Users) DeleteUser() error {
	if result := database.DB.Delete(&u); result.Error != nil {
		log.Printf("Failed to delete user with ID: %d | Error: %s", u.ID, result.Error.Error())
		return result.Error
	}
	return nil
}

func ValidateUserExistsWithEmailOrUsername(params UserCheckParams) (*Users, string, bool) {
	var count int64
	var dbUser Users
	msg := ""
	exists := false
	query := database.DB.Model(&Users{}).Where("username = ? OR email = ?", params.Username, params.Email)
	if params.UserId > 0 {
		query = query.Not("id = ?", params.UserId)
	}

	if err := query.First(&dbUser).Error; err == nil {
		log.Printf("Found User with Username: %s | Email: %s", params.Username, params.Email)
		msg = "User exists with the given attribute(s)"
		exists = true
	}

	if err := query.Count(&count).Error; err == nil {
		log.Printf("Found %d User with Username: %s | Email: %s", count, params.Username, params.Email)
		msg = "Several User Exists with the given attribute(s)"
		exists = count > 0
	}

	return &dbUser, msg, exists
}

func (u *Users) ValidateUserExists() (string, bool) {
	userParams := UserCheckParams{
		UserId:   u.ID,
		Username: u.Username,
		Email:    u.Email,
	}
	_, msg, exists := ValidateUserExistsWithEmailOrUsername(userParams)
	return msg, exists
}

func (u *Users) Save() error {
	if u.ID == 0 {
		if result := database.DB.Create(&u); result.Error != nil {
			log.Printf("Failed to create user: %s", result.Error.Error())
			return result.Error
		}
	} else {
		if result := database.DB.Save(&u); result.Error != nil {
			log.Printf("Failed to save user: %s", result.Error.Error())
			return result.Error
		}
	}
	return nil
}

func (u *Users) GeneratePasswordHash() (error, bool) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to Generate Password, Error: %s", err.Error())
		return err, false
	}
	u.Password = string(hashedPassword)
	return nil, true
}

func (user *Users) ValidatePasswordHash(password string) (string, bool) {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Printf("comparing token got error %s", err.Error())
		return "Invalid credentials", false
	}
	return "", true
}

func (userToUpdate *Users) UpdateUser(updateData interface{}, omitFields ...string) (*Users, error) {
	if result := database.DB.Model(userToUpdate).Omit(omitFields...).Updates(updateData); result.Error != nil {
		log.Printf("Failed to update User: %s | Error: %s", updateData, result.Error.Error())
		return nil, result.Error
	}
	return userToUpdate, nil
}
