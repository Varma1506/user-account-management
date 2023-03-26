package services

import (
	"fmt"
	"regexp"

	model "github.com/Varma1506/user-account-management/model"
)

func SignupRequestValidator(req model.SignupRequest) error {
	if req.Firstname == "" || req.Lastname == "" {
		return fmt.Errorf("Firstname/Lastname cannot be empty")
	} else if req.Password == "" || req.Email == "" {
		return fmt.Errorf("Password/Email cannot be empty")
	} else if !isValidEmail(req.Email) {
		return fmt.Errorf("Invalid Email")
	} else if len(req.Password) < 8 {
		return fmt.Errorf("password must be 8 or more characters long")
	} else if req.ConfirmPassword != req.Password {
		return fmt.Errorf("password is not matching with confirm password")
	} else if req.Username == "" {
		return fmt.Errorf("username cannot be empty")
	} else if len(req.Password) < 8 {
		return fmt.Errorf("password must be 8 or more characters long")
	}
	return nil
}

func ChangePassowrdRequestValidator(req model.ChangePassowrdRequest) error {
	if req.Username == "" {
		return fmt.Errorf("username cannot be empty")
	} else if req.Currentpass == "" {
		return fmt.Errorf("currentpass cannot be empty")
	} else if req.NewPass == req.Currentpass {
		return fmt.Errorf("new password shouldn't match with Old password")
	}
	return nil
}

func LoginRequestValidator(req model.LoginRequest) error {
	if req.Username == "" {
		return fmt.Errorf("invalid/missing username")
	} else if req.Password == "" {
		return fmt.Errorf("invalid/missing password")
	}
	return nil
}

func isValidEmail(email string) bool {
	// Regular expression for validating an email address
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regular expression
	regex := regexp.MustCompile(pattern)

	// Check if the email matches the pattern
	return regex.MatchString(email)
}
