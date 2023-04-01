package token

import (
	"fmt"
	"time"

	model "github.com/Varma1506/user-account-management/model"
	"github.com/dgrijalva/jwt-go"
)

// get secret key
func getSecretKey() []byte {
	return []byte("userauthtestproject")
}

// generate token with username encoded
func GenerateToken(username string) string {
	secretKey := getSecretKey()
	//Configure header and payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 1).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "error"
	}
	return tokenString
}

// Get data from the token
func GetDataFromToken(tokenString string) (*model.TokenClaim, error) {
	claims := &model.TokenClaim{}
	secretKey := getSecretKey()

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
