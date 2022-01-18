package dbfunctions

import (
	"fmt"
	"shopifyproductionengineertest/inventory"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)
func ListAllProducts(client *mongo.Client) (inventory.ProductList, error) {
	

	coll := client.Database(DbName).Collection(ProductCollectionName)
	cursor, err := coll.Find(ctx, bson.M{})
	
	if err != nil {
		return inventory.ProductList{}, err
	}

	var products inventory.ProductList
	for cursor.Next(ctx) {
		var p inventory.Product
		err := cursor.Decode(&p)
		if err != nil {
			return inventory.ProductList{}, err
		}
		fmt.Println(p)

		products.Products = append(products.Products, p)
	}

	

	return products, nil
}