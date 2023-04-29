package helpers

import (
	"log"
	"strconv"
	"time"

	"github.com/fahad-md-kamal/fiber-blogs/configs"
	"github.com/fahad-md-kamal/fiber-blogs/users/dtos"
	"github.com/fahad-md-kamal/fiber-blogs/users/models"
	"github.com/golang-jwt/jwt"
)

func GenerateJwtToken(user *models.Users) (string, bool) {
	tokenLifetime, err := strconv.ParseInt(configs.ENVs.TokenLifeTime, 10, 0)
	if err != nil {
		log.Printf("Failed to read token lifetime environment variable: %s", err.Error())
		return "Failed to read token lifetime environment variable", false
	}
	claims := &dtos.TokenClaimPayload{
		ID:       user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(tokenLifetime) * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(configs.ENVs.JwtSecretKey))
	if err != nil {
		log.Printf("Internal server error while attaching signe to the token: %s", err.Error())
		return "Error for generating signed token", false
	}

	return tokenString, true
}
