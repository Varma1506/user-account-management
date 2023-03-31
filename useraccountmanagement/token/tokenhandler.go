package token

import (
	"fmt"

	model "github.com/Varma1506/user-account-management/model"
	"github.com/dgrijalva/jwt-go"
)

// generate token with username encoded
func GenerateToken(username string) string {

	secretString := "userauthtestproject"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})

	tokenString, err := token.SignedString([]byte(secretString))
	if err != nil {
		return "error"
	}
	return tokenString
}

// Validate the token
func ValidateToken(tokenString string) (model.TokenClaim, error) {
	claims := &model.TokenClaim{}
	secretKey := "userauthtestproject"

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(*model.TokenClaim); !ok && !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}

	return claims, nil
}
