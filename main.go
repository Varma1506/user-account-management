package main

import (
	"fmt"
	"net/http"

	"./routes"
)

func main() {

	http.HandleFunc("/sign-up", routes.Signup)
	http.HandleFunc("/change-password", routes.ChangePassowrd)
	http.HandleFunc("/delete-account/", routes.DeleteAccount)
	http.HandleFunc("/login", routes.Login)
	fmt.Println("SUccess")
	http.ListenAndServe(":8000", nil)
}
