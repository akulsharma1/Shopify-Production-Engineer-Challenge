package dbfunctions

import (
	"fmt"
	"shopifyproductionengineertest/inventory"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)
func ListAllProducts(client *mongo.Client) error {
	

	// var v2 []inventory.Product
	coll := client.Database(DbName).Collection(ProductCollectionName)
	// _, err := coll.Find(ctx, bson.M{}).All(ctx, &p)
	cursor, err := coll.Find(ctx, bson.M{})
	
	if err != nil {
		fmt.Println("----")
		fmt.Println(err)
		fmt.Println("----")
	}

	for cursor.Next(ctx) {
		var p inventory.Product
		err := cursor.Decode(&p)
		if err != nil {
			fmt.Println("***")
			fmt.Println(err)
			fmt.Println("***")
		}
		fmt.Println(p)

		//v = append(v, p)
	}

	// log.Println(v)

	

	return nil
}