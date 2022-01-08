package main

import (
	"shopifyproductionengineertest/inventory"
	"shopifyproductionengineertest/managerfunctions"
)

func main() {
	managerfunctions.ReadProducts()
	managerfunctions.CreateInventoryItem("test", []inventory.Variant{inventory.Variant{VariantID: 1, VariantName: "test1", Size: "test", Image: "test", Quantity: 1, Price: 1.0}})
}