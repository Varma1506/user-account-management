package services

import (
	"encoding/json"
	"net/http"

	"../model"
)

func BuildResponse(w http.ResponseWriter, code int, message string, data []model.User) {
	var buildRequestResponse model.Response
	buildRequestResponse.Status = code
	buildRequestResponse.Message = message
	buildRequestResponse.Data = data

	w.WriteHeader(code)
	json.NewEncoder(w).Encode(buildRequestResponse)
}