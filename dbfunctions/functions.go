package dbfunctions

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoURL = "mongodb+srv://user1:testpass@cluster0.mlnm2.mongodb.net/Product-List?retryWrites=true&w=majority"
	DbName = "Product-List"
	ProductCollectionName = "Products"
)
var ctx = context.TODO()

func Init() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(MongoURL)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return client, nil
}

