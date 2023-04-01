package auth

import (
	"fmt"
	"net/http"

	model "github.com/Varma1506/user-account-management/model"
	services "github.com/Varma1506/user-account-management/services"
	"github.com/dgrijalva/jwt-go"
)

func getSecretKey() []byte {
	return []byte("userauthtestproject")
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Secret Keys
		secretKey := getSecretKey()
		tokenString := r.Header.Get("Authorization")

		//Parse token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				services.BuildResponse(w, http.StatusUnauthorized, "unexpected signing method", []model.User{})
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secretKey, nil
		})
		if err != nil {
			services.BuildResponse(w, http.StatusUnauthorized, err.Error(), []model.User{})
			return
		}

		if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
			services.BuildResponse(w, http.StatusUnauthorized, "Invalid Token", []model.User{})
			return
		}
		next(w, r)
	})
}
