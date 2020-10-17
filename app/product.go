package coffeego

import (
	helper "coffeego/helper"
	"encoding/json"
	"net/http"
)

//ProductListRequest Should be exported
func ProductListRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type response struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	var resp response
	resp.ID = 1
	resp.Title = "kucing"
	resp.Body = "kucingkucingkucingkucing kucingkucingkucing  kucingkucing"

	output, err := json.Marshal(resp)
	if err != nil {
		helper.ReponseOnServerError(w)
		return
	}
	w.Write(output)
}
