package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	model "github.com/Varma1506/user-account-management/model"
)

// Build a Response
func BuildResponse(w http.ResponseWriter, code int, message string, data []model.User) {
	var buildRequestResponse model.Response
	buildRequestResponse.Status = code
	buildRequestResponse.Message = message
	buildRequestResponse.Data = data

	w.WriteHeader(code)
	json.NewEncoder(w).Encode(buildRequestResponse)
}

// Extract body from request
func ExtractReqBody(r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("Error in parsing data")
	}
	return body, nil
}
