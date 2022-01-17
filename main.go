package main

import (
	"shopifyproductionengineertest/dbfunctions"

	// "shopifyproductionengineertest/managerfunctions"
	"fmt"
)

func main() {
	client, err := dbfunctions.Init()
	if err != nil {
		fmt.Println(err)
	}

	// err = dbfunctions.AddProduct(client, inventory.Product{
	// 	ProductName: "Test Product",
	// 	ProductID:   "testproduct",
	// 	Variants: []inventory.Variant{
	// 		inventory.Variant{
	// 			VariantID: "testvariant",
	// 			VariantName: "Test Variant",
	// 			Size: "Test Size",
	// 			Image: "Test Image",
	// 			Quantity:  10,
	// 			Price:     10.00,
	// 		},
	// 	},
	// })

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// err = dbfunctions.AddVariant(client, "testproduct", inventory.Variant{
	// 	VariantID: "testvariant2",
	// 	VariantName: "Test Variant",
	// 	Size: "Test Size",
	// 	Image: "Test Image",
	// 	Quantity:  10,
	// 	Price:     10.00,
	// })

	// err = dbfunctions.AddVariant(client, "testproduct", inventory.Variant{
	// 	VariantID: "testvariant2",
	// 	VariantName: "Test Variant",
	// 	Size: "Test Size",
	// 	Image: "Test Image",
	// 	Quantity:  10,
	// 	Price:     10.00,
	// })

	// err = dbfunctions.EditVariant(client, "testproduct", inventory.Variant{
	// 	VariantID: "testvariant2",
	// 	VariantName: "Test Variant edit",
	// 	Size: "Test Size",
	// 	Image: "Test Image",
	// 	Quantity:  10,
	// 	Price:     10.00,
	// })
		
		dbfunctions.ListAllProducts(client)
	// })

	if err != nil {
		fmt.Println(err)
	}
}