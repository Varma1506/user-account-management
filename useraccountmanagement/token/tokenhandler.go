package token

import (
	"fmt"
	"time"

	model "github.com/Varma1506/user-account-management/model"
	"github.com/dgrijalva/jwt-go"
)

// generate token with username encoded
func GenerateToken(username string) string {

	secretString := "userauthtestproject"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretString))
	if err != nil {
		return "error"
	}
	return tokenString
}

// Validate the token
func ValidateToken(tokenString string) (*model.TokenClaim, error) {
	claims := &model.TokenClaim{}
	secretString := "userauthtestproject"
	secretKey := []byte(secretString)

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return &model.TokenClaim{}, fmt.Errorf(err.Error())
	}

	if token.Valid {
		return claims, nil
	}
	return &model.TokenClaim{}, fmt.Errorf("Invalid Token")
}
