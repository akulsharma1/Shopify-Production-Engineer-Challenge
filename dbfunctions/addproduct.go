package dbfunctions

import (
	"fmt"
	"shopifyproductionengineertest/inventory"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddProduct(client *mongo.Client, product inventory.Product) error {
	coll := client.Database(DbName).Collection(ProductCollectionName)
	_, err := coll.InsertOne(ctx, product)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func AddVariant(client *mongo.Client, productID string, variant inventory.Variant) error {
	coll := client.Database(DbName).Collection(ProductCollectionName)
	_, err := coll.UpdateOne(ctx, bson.M{"productID": productID}, bson.M{"$push": bson.M{"variants": variant}})
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
