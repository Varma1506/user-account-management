package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	dbconfig "github.com/Varma1506/user-account-management/dbconfig"
	model "github.com/Varma1506/user-account-management/model"
	services "github.com/Varma1506/user-account-management/services"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var data []model.User
	if r.Method == http.MethodPost {
		//Connect to the DB
		db := dbconfig.Connect()
		defer db.Close()

		//Process the request body to access the data
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			services.BuildResponse(w, http.StatusBadRequest, "Error in parsing body from request", data)
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
		userFromDB := GetUserRecord(user.Username)
		if userFromDB.Username != "" {
			services.BuildResponse(w, http.StatusBadRequest, "User with given username already exists, try another one", data)
			return
		}

		_, err = db.Exec("INSERT INTO `go-devs` (`full name`,username,email,password) VALUES(?,?,?,?)", user.Firstname+" "+user.Lastname, user.Username, user.Email, user.Password)
		if err != nil {
			services.BuildResponse(w, http.StatusInternalServerError, "Error in adding data to DB", data)
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
		//establish db connection
		db := dbconfig.Connect()
		defer db.Close()

		//Extract the request body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			services.BuildResponse(w, http.StatusBadRequest, "Error in getting body from request", data)
			return
		}

		var pass model.ChangePassowrdRequest
		err = json.Unmarshal(body, &pass)
		if err != nil {
			services.BuildResponse(w, http.StatusBadRequest, "Error parsing request body", data)
			return
		}

		//validate the request
		err = services.ChangePassowrdRequestValidator(pass)
		if err != nil {
			services.BuildResponse(w, http.StatusBadRequest, err.Error(), data)
			return
		}

		//Check for password in DB
		userFromDB := GetUserRecord(pass.Username)
		if userFromDB.Username == "" {
			services.BuildResponse(w, http.StatusBadRequest, "User doesn't exist with the given username", data)
			return
		} else if userFromDB.Password != pass.Currentpass {
			services.BuildResponse(w, http.StatusBadRequest, "Current password doesn't match with password in the record in DB", data)
			return
		}

		//Update password
		_, err = db.Exec("UPDATE `go-devs` SET password=? WHERE username=?", pass.NewPass, pass.Username)
		if err != nil {
			services.BuildResponse(w, http.StatusInternalServerError, "Error in Updating DB", data)
			return
		}

		//success response
		services.BuildResponse(w, http.StatusOK, "Updated Succesfully", data)
	} else {
		services.BuildResponse(w, http.StatusMethodNotAllowed, "Not a PUT request", data)
	}
}

// Delete a account
func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is hitted")
	var data []model.User
	if r.Method == http.MethodDelete {
		//establish db connection
		db := dbconfig.Connect()
		defer db.Close()

		//Extract username from URI path
		path := r.URL.Path
		segments := strings.Split(path, "/")
		username := segments[2]

		//Check if the record exists
		userFromDB := GetUserRecord(username)
		if userFromDB.Username == "" {
			services.BuildResponse(w, http.StatusBadRequest, "User doesn't exist with the given username", data)
			return
		}

		//Delete record from DB
		_, err := db.Exec("DELETE FROM `go-devs` where username=?", username)
		if err != nil {
			services.BuildResponse(w, http.StatusInternalServerError, "Error in deleting record in DB", data)
			return
		}

		//send success response
		w.WriteHeader(http.StatusNoContent)

	} else {
		services.BuildResponse(w, http.StatusMethodNotAllowed, "Not a DELETE Request", data)
	}
}

// login function
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is hit : ", r.Method)
	var data []model.User
	if r.Method == http.MethodPost {
		//Connect to DB
		db := dbconfig.Connect()
		defer db.Close()

		//Extract body from Req
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			services.BuildResponse(w, http.StatusBadRequest, "Error in getting body from request", data)
			return
		}

		var login model.LoginRequest
		err = json.Unmarshal(body, &login)
		if err != nil {
			services.BuildResponse(w, http.StatusBadRequest, "Error parsing request body", data)
			return
		}
		//Validate Input
		err = services.LoginRequestValidator(login)
		if err != nil {
			services.BuildResponse(w, http.StatusBadRequest, err.Error(), data)
		}

		//Check if user exists in DB
		userFromDB := GetUserRecord(login.Username)
		if userFromDB.Username == "" {
			services.BuildResponse(w, http.StatusBadRequest, "User doesn't exist with the given username", data)
			return
		}
		//check if password matches with record in DB
		if userFromDB.Password != login.Password {
			services.BuildResponse(w, http.StatusBadRequest, "Wrong Password please try again", data)
			return
		}

		//Success Response
		services.BuildResponse(w, http.StatusOK, "Login successful", data)

	} else {
		services.BuildResponse(w, http.StatusMethodNotAllowed, "Not a POST request", data)
	}
}

// To get record from DB with given username f
func GetUserRecord(username string) model.User {
	db := dbconfig.Connect()
	defer db.Close()

	var user model.User
	err := db.QueryRow("SELECT userid,`full name`,username,email,password FROM `go-devs` WHERE username=?", username).Scan(&user.Id, &user.Fullname, &user.Username, &user.Email, &user.Password)
	if err != nil {
		fmt.Println("This is Error")
	}
	return user
}
