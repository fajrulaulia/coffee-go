package coffeego

import (
	"encoding/json"
	"net/http"

	helper "coffeego/helper"
)

//UserListRequest Should be exported
func UserListRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type response struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"Email"`
	}
	var resp response

	resp.ID = 1
	resp.Username = "Fajrul Aulia"
	resp.Email = "Auliafajrul7@gmail.com"

	output, err := json.Marshal(resp)
	if err != nil {
		helper.ReponseOnServerError(w)
		return
	}
	w.Write(output)
}

//CurrentUserRequest Should be exported
func CurrentUserRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type response struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"Email"`
	}
	var resp response

	resp.ID = 1
	resp.Username = "Fajrul Aulia"
	resp.Email = "Auliafajrul7@gmail.com"

	output, err := json.Marshal(resp)
	if err != nil {
		helper.ReponseOnServerError(w)
		return
	}
	w.Write(output)
}
