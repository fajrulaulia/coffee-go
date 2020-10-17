package coffeego

import (
	helper "coffeego/helper"
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

//ProductRead Should be exported
func ProductRead(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type response struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	var result bson.M

	vars := mux.Vars(r)
	//5f8ab88aa528413066e65fa5

	docID, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		helper.ReponseOnServerError(w)
		return
	}

	err = helper.Collection("products").FindOne(context.TODO(), bson.M{"_id": docID}).Decode(&result)
	if err != nil {
		helper.ReponseOnServerError(w)
		return
	}

	var resp response
	resp.ID = result["_id"].(primitive.ObjectID).Hex()
	resp.Name = result["name"].(string)
	resp.Description = result["description"].(string)

	output, err := json.Marshal(resp)
	if err != nil {
		helper.ReponseOnServerError(w)
		return
	}
	w.Write(output)
}
