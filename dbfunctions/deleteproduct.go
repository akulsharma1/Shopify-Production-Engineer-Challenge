package dbfunctions

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteProduct(client *mongo.Client, productID string) error {
	coll := client.Database(DbName).Collection(ProductCollectionName)
	_, err := coll.DeleteOne(ctx, bson.M{"productID": productID})
	if err != nil {
		return err
	}

	return nil
}

func DeleteVariant(client *mongo.Client, productID string, variantID string) error {
	coll := client.Database(DbName).Collection(ProductCollectionName)
	filter := bson.M {
		"productid": productID,
	}
	_, err := coll.UpdateOne(ctx, filter, bson.M{"$pull": bson.M{"variants": bson.M{"variantid": variantID}}})
	if err != nil {
		return err
	}

	return nil
}