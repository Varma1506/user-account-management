package main

import (
	"fmt"
	"net/http"

	"./routes"
)

func main() {

	http.HandleFunc("/signup", routes.Signup)
	http.HandleFunc("/changepassword", routes.ChangePassowrd)
	http.HandleFunc("/deleteaccount/", routes.DeleteAccount)
	http.HandleFunc("/login", routes.Login)
	fmt.Println("SUccess")
	http.ListenAndServe(":8000", nil)
}
