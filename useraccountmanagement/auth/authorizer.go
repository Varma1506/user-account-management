package auth

import (
	"net/http"

	model "github.com/Varma1506/user-account-management/model"
	services "github.com/Varma1506/user-account-management/services"
	"github.com/dgrijalva/jwt-go"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := &model.TokenClaim{}
		secretString := "userauthtestproject"
		secretKey := []byte(secretString)

		tokenString := r.Header.Get("Authorization")
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil {
			services.BuildResponse(w, http.StatusUnauthorized, err.Error(), []model.User)
			return
		}

		if !token.Valid {
			services.BuildResponse(w, http.StatusUnauthorized, "Invalid Token", []model.User)
			return
		}
		next(w, r)
	})
}
