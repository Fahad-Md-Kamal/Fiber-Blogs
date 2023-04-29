package dtos

import (
	"time"

	"github.com/fahad-md-kamal/fiber-blogs/users/models"
	"github.com/fahad-md-kamal/fiber-blogs/utils"
)

type UserCreateDto struct {
	Username string `json:"username" validate:"required,min=4,max=50"`
	Email    string `json:"email" validate:"required,email,min=8,max=100"`
	Password string `json:"password" validate:"required,min=6"`
}

func (data *UserCreateDto) ValidateUserCreateDto() ([]*utils.ErrorResponse, bool) {
	errors := utils.ValidateStruct(data)
	return errors, len(errors) == 0
}

func (data *UserCreateDto) ParseFromDto() (user *models.Users) {
	user = &models.Users{}
	user.Username = data.Username
	user.Email = data.Email
	user.Password = data.Password
	return user
}

type UserUpdateDto struct {
	Email       string `json:"email" validate:"omitempty,email,min=8,max=100"`
	IsSuperuser *bool  `json:"is_superuser" validate:"omitempty"`
	IsActive    *bool  `json:"is_active" validate:"omitempty"`
}

func (data *UserUpdateDto) ValidateUserUpdateDto() ([]*utils.ErrorResponse, bool) {
	errors := utils.ValidateStruct(data)
	return errors, len(errors) == 0
}

type UserResponseDto struct {
	Id          uint      `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	IsActive    bool      `json:"is_active"`
	IsSuperuser bool      `json:"is_superuser"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (udto *UserResponseDto) ParseToResponseDto(user *models.Users) {
	udto.Id = user.ID
	udto.Username = user.Username
	udto.Email = user.Email
	udto.IsSuperuser = user.IsSuperuser
	udto.IsActive = user.IsActive
	udto.CreatedAt = user.CreatedAt
	udto.UpdatedAt = user.UpdatedAt
}

func ParseUsersListToResponseDto(users *[]models.Users) []UserResponseDto {
	usersList := []UserResponseDto{}
	for _, user := range *users {
		usersList = append(usersList, UserResponseDto{
			Id:          user.ID,
			Username:    user.Username,
			Email:       user.Email,
			IsActive:    user.IsActive,
			IsSuperuser: user.IsSuperuser,
			CreatedAt:   user.CreatedAt,
			UpdatedAt:   user.UpdatedAt,
		})
	}
	return usersList
}

func ParseUserToResponseDto(user *models.Users) *UserResponseDto {
	userDto := UserResponseDto{
		Id:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		IsActive:    user.IsActive,
		IsSuperuser: user.IsSuperuser,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
	return &userDto
}
