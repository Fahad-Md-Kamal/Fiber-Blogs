package dtos

import (
	"github.com/fahad-md-kamal/fiber-blogs/users/models"
	"github.com/fahad-md-kamal/fiber-blogs/utils"
	"github.com/golang-jwt/jwt"
)

type TokenClaimPayload struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type LoginRequestDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (data *LoginRequestDto) ValidateLoginRequestDto() ([]*utils.ErrorResponse, bool) {
	errors := utils.ValidateStruct(data)
	return errors, len(errors) == 0
}

type LoginResponseDto struct {
	UserID      uint   `json:"userId"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Token       string `json:"token"`
	IsSuperuser bool   `json:"isSuperUser"`
}

func ParseToLoginResponseDto(token string, u *models.Users) *LoginResponseDto {
	loginResponseDto := LoginResponseDto{
		UserID:   u.ID,
		Username: u.Username,
		Email:    u.Email,
		Token:    token,
	}

	return &loginResponseDto
}
