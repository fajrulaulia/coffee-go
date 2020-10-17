package coffeego

import (
	core "coffeego/helper/core"

	"go.mongodb.org/mongo-driver/mongo"
)

//Collection should be exported
func Collection(name string) *mongo.Collection {
	client := core.InitDB()
	return client.Database("coffeedb").Collection(name)
}
