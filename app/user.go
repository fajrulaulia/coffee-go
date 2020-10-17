package coffeego

import (
	"encoding/json"
	"fmt"
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
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("{\"email\":\"%v\",\"status_code\":\"%v\"}", "email", 200)))
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
