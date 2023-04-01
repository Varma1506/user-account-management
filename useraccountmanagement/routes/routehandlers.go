package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	dbconfig "github.com/Varma1506/user-account-management/dbconfig"
	model "github.com/Varma1506/user-account-management/model"
	services "github.com/Varma1506/user-account-management/services"
	token "github.com/Varma1506/user-account-management/token"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var data []model.User
	if r.Method == http.MethodPost {

		//Process the request body to access the data
		body, err := services.ExtractReqBody(r)
		if err != nil {
			services.BuildResponse(w, http.StatusBadRequest, err.Error(), data)
			return
		}

		var user model.SignupRequest
		err = json.Unmarshal(body, &user)
		if err != nil {
			services.BuildResponse(w, http.StatusBadRequest, "Error parsing request body", data)
			return
		}

		//validate Request Body
		err = services.SignupRequestValidator(user)

		if err != nil {
			services.BuildResponse(w, http.StatusBadRequest, err.Error(), data)
			return
		}
		//Check if Username exists
		userFromDB := dbconfig.GetUserRecord(user.Username)
		if userFromDB.Username != "" {
			services.BuildResponse(w, http.StatusBadRequest, "User with given username already exists, try another one", data)
			return
		}

		err = dbconfig.InsertUserIntoDB(user)
		if err != nil {
			services.BuildResponse(w, http.StatusInternalServerError, err.Error(), data)
			return
		}
		//Send success Response
		services.BuildResponse(w, http.StatusCreated, "Created a new user", data)
	} else {
		services.BuildResponse(w, http.StatusMethodNotAllowed, "Not a POST request", data)
	}

}

func ChangePassowrd(w http.ResponseWriter, r *http.Request) {
	var data []model.User
	if r.Method == http.MethodPut {
		//Validate Token
		tokenString := r.Header.Get("Authorization")
		tokenValidatorResponse, err := token.ValidateToken(tokenString)
		if err != nil {
			services.BuildResponse(w, http.StatusUnauthorized, err.Error(), data)
			return
		}

		//Extract the request body
		body, err := services.ExtractReqBody(r)
		if err != nil {
			services.BuildResponse(w, http.StatusBadRequest, err.Error(), data)
			return
		}

		var pass model.ChangePassowrdRequest
		err = json.Unmarshal(body, &pass)
		if err != nil {
			services.BuildResponse(w, http.StatusInternalServerError, "Error parsing request body", data)
			return
		}

		//Check if username for the password change request matches with Logged User
		if pass.Username != tokenValidatorResponse.Username {
			services.BuildResponse(w, http.StatusUnauthorized, "User is not logged in", data)
			return
		}

		//validate the request
		err = services.ChangePassowrdRequestValidator(pass)
		if err != nil {
			services.BuildResponse(w, http.StatusBadRequest, err.Error(), data)
			return
		}

		//Check for password in DB
		userFromDB := dbconfig.GetUserRecord(pass.Username)
		if userFromDB.Username == "" {
			services.BuildResponse(w, http.StatusBadRequest, "User doesn't exist with the given username", data)
			return
		} else if userFromDB.Password != pass.Currentpass {
			services.BuildResponse(w, http.StatusBadRequest, "Current password doesn't match with password in the record in DB", data)
			return
		}

		//Update password
		err = dbconfig.UpdateUserPassword(pass.Username, pass.NewPass)
		if err != nil {
			services.BuildResponse(w, http.StatusInternalServerError, err.Error(), data)
		}

		//success response
		services.BuildResponse(w, http.StatusOK, "Updated Succesfully", data)
	} else {
		services.BuildResponse(w, http.StatusMethodNotAllowed, "Not a PUT request", data)
	}
}

// Delete a account
func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	var data []model.User
	if r.Method == http.MethodDelete {
		//Validate Token
		tokenString := r.Header.Get("Authorization")
		tokenValidatorResponse, err := token.ValidateToken(tokenString)
		if err != nil {
			services.BuildResponse(w, http.StatusUnauthorized, err.Error(), data)
			return
		}

		//Extract username from URI path
		path := r.URL.Path
		segments := strings.Split(path, "/")
		username := segments[2]

		//Check if username for the password change request matches with Logged User
		if username != tokenValidatorResponse.Username {
			services.BuildResponse(w, http.StatusUnauthorized, "User is not logged in", data)
			return
		}

		//Check if the record exists
		userFromDB := dbconfig.GetUserRecord(username)
		if userFromDB.Username == "" {
			services.BuildResponse(w, http.StatusBadRequest, "User doesn't exist with the given username", data)
			return
		}

		//Delete record from DB
		err = dbconfig.DeleteUserRecord(username)
		if err != nil {
			services.BuildResponse(w, http.StatusInternalServerError, err.Error(), data)
		}

		//send success response
		w.WriteHeader(http.StatusNoContent)

	} else {
		services.BuildResponse(w, http.StatusMethodNotAllowed, "Not a DELETE Request", data)
	}
}

// login function
func Login(w http.ResponseWriter, r *http.Request) {
	var data []model.User
	if r.Method == http.MethodPost {
		var login model.LoginRequest
		//Extract body from Req
		body, err := services.ExtractReqBody(r)
		if err != nil {
			services.BuildResponse(w, http.StatusBadRequest, err.Error(), data)
			return
		}
		err = json.Unmarshal(body, &login)

		if err != nil {
			services.BuildResponse(w, http.StatusInternalServerError, "Error parsing request body", data)
			return
		}
		//Validate Input
		err = services.LoginRequestValidator(login)
		if err != nil {
			services.BuildResponse(w, http.StatusBadRequest, err.Error(), data)
		}

		//Check if user exists in DB
		userFromDB := dbconfig.GetUserRecord(login.Username)
		if userFromDB.Username == "" {
			services.BuildResponse(w, http.StatusBadRequest, "User doesn't exist with the given username", data)
			return
		}
		//check if password matches with record in DB
		if userFromDB.Password != login.Password {
			services.BuildResponse(w, http.StatusBadRequest, "Wrong Password please try again", data)
			return
		}

		//Generate Token
		tokenString := token.GenerateToken(login.Username)

		//Success Response
		services.BuildResponse(w, http.StatusOK, tokenString, data)

	} else {
		services.BuildResponse(w, http.StatusMethodNotAllowed, "Not a POST request", data)
	}
}
