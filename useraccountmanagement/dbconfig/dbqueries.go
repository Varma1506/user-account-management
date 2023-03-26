package dbconfig

import (
	"fmt"

	model "github.com/varma1506/user-account-management/models"
)

func InsertUserIntoDB(user model.User) error {
	//creating db session
	db := Connect()
	defer db.Close()

	//execute query
	_, err := db.Exec("INSERT INTO `go-devs` (`full name`,username,email,password) VALUES(?,?,?,?)", user.Firstname+" "+user.Lastname, user.Username, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("Error in inserting data into db")
	}
	return nil
}

// To get record from DB with given username f
func GetUserRecord(username string) model.User {
	//Create instance of DB
	db := Connect()
	defer db.Close()

	var user model.User
	err := db.QueryRow("SELECT userid,`full name`,username,email,password FROM `go-devs` WHERE username=?", username).Scan(&user.Id, &user.Fullname, &user.Username, &user.Email, &user.Password)
	if err != nil {
		fmt.Println("This is Error")
	}
	return user
}

// To delete a record in DB
func DeleteUserRecord(username string) error {
	//Create instance of DB
	db := Connect()
	defer db.Close()

	_, err := db.Exec("DELETE FROM `go-devs` where username=?", username)
	if err != nil {
		return fmt.Errorf("Error in Deleting Record in DB")
	}
	return nil
}

// To update User Password
func UpdateUserPassword(username string, password string) error {
	//Create instance of DB
	db := Connect()
	defer db.Close()

	_, err := db.Exec("UPDATE `go-devs` SET password=? WHERE username=?", password, username)
	if err != nil {
		return fmt.Errorf("Error in updating Password in DB")
	}
	return nil
}
