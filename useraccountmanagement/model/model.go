package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	Id       int    `json:"userid"`
	Fullname string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type Response struct {
	Status  int    `json:"Status"`
	Message string `json:"Message"`
	Data    []User `json:"Data"`
}

type SignupRequest struct {
	Firstname       string `json:"first_name"`
	Lastname        string `json:"last_name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Username        string `json:"username"`
	ConfirmPassword string `json:"cnfpass"`
}

type ChangePassowrdRequest struct {
	Username    string `json:"username"`
	Currentpass string `json:"current_password"`
	NewPass     string `json:"new_password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
