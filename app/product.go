package coffeego

import (
	helper "coffeego/helper"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type productResponse struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	SellPrice   float64            `json:"sell_price" bson:"sell_price"`
	BuyPrice    float64            `json:"buy_price" bson:"buy_price"`
	CreatedAt   primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt   primitive.DateTime `json:"updated_at" bson:"updated_at"`
}

//ProductListRequest Should be exported
func ProductListRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var resp []productResponse
	var item productResponse
	findOptions := options.Find()

	var page int64 = 0
	var limit int64 = 0
	if r.URL.Query().Get("page") != "" && r.URL.Query().Get("limit") != "" {
		page, err := strconv.ParseInt(r.URL.Query().Get("page"), 10, 32)
		if err != nil {
			helper.ReponseOnServerError(w)
			return
		}
		limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 32)
		if err != nil {
			helper.ReponseOnServerError(w)
			return
		}
		if page == 1 {
			findOptions.SetSkip(0)
			findOptions.SetLimit(limit)
		} else {
			findOptions.SetSkip((page - 1) * limit)
			findOptions.SetLimit(limit)
		}
	}
	cur, err := helper.Collection("products").Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		helper.ReponseOnServerError(w)
		return
	}

	type response struct {
		Total int               `json:"total"`
		Limit int64             `json:"limit"`
		Page  int64             `json:"page"`
		Data  []productResponse `json:"data"`
	}
	var headerResponse response

	for cur.Next(context.TODO()) {
		err := cur.Decode(&item)
		if err != nil {
			helper.ReponseOnServerError(w)
			return
		}
		resp = append(resp, item)
	}
	headerResponse.Data = resp
	headerResponse.Limit = limit
	headerResponse.Page = page
	headerResponse.Total = len(resp)
	output, err := json.Marshal(headerResponse)
	if err != nil {
		return
	}
	w.Write(output)
}

//ProductRead Should be exported
func ProductRead(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	docID, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	if err != nil {
		helper.ReponseOnNotFound(w)
		return
	}
	var result bson.M
	err = helper.Collection("products").FindOne(context.TODO(), bson.M{"_id": docID}).Decode(&result)
	if err != nil {
		helper.ReponseOnNotFound(w)
		return
	}
	var resp productResponse
	resp.ID = result["_id"].(primitive.ObjectID)
	resp.Name = result["name"].(string)
	resp.Description = result["description"].(string)
	resp.SellPrice = result["sell_price"].(float64)
	resp.BuyPrice = result["buy_price"].(float64)
	resp.CreatedAt = result["created_at"].(primitive.DateTime)
	resp.UpdatedAt = result["updated_at"].(primitive.DateTime)
	output, err := json.Marshal(resp)
	if err != nil {
		helper.ReponseOnServerError(w)
		return
	}
	w.Write(output)
}
