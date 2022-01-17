package dbfunctions

import (
	"shopifyproductionengineertest/inventory"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func EditProduct(client *mongo.Client, product inventory.Product) error {
	coll := client.Database(DbName).Collection(ProductCollectionName)
	_, err := coll.UpdateOne(ctx, bson.M{"productID": product.ProductID}, bson.M{"$set": product})

	if err != nil {
		return err
	}

	return nil
}

func EditVariant(client *mongo.Client, productID string, variant inventory.Variant) error {
	coll := client.Database(DbName).Collection(ProductCollectionName)
	opts := options.Update().SetUpsert(true)
	filter := bson.M {
		"productID": productID,
		"variants.variantID": variant.VariantID,
	}
	_, err := coll.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"variants": variant}}, opts)
	if err != nil {
		return err
	}

	return nil
}