package main

import (
	"fmt"
	"net/http"

	auth "github.com/Varma1506/user-account-management/auth"
	routes "github.com/Varma1506/user-account-management/routes"
)

func main() {

	http.HandleFunc("/signup", routes.Signup)
	http.HandleFunc("/changepassword", auth.Authenticate(routes.ChangePassowrd))
	http.HandleFunc("/deleteaccount/", auth.Authenticate(routes.DeleteAccount))
	http.HandleFunc("/login", routes.Login)
	fmt.Println("SUccess")
	http.ListenAndServe(":8000", nil)
}
