package managerfunctions

import (
	"fmt"
	"shopifyproductionengineertest/inventory"
)
func DeleteProduct(id string) error {
	prodList, err := ReadProducts()
	if err != nil {
		fmt.Println(err)
		return err
	}

	currProdList := inventory.ProductList{}
	for i := range prodList.Products {
		if prodList.Products[i].ProductID != id {
			currProdList.Products = append(currProdList.Products, prodList.Products[i])
		}
	}

	return nil
}

func DeleteInventoryVariant(productID, variantID string) error {
	prodList, err := ReadProducts()
	if err != nil {
		fmt.Println(err)
		return err
	}

	currProdList := inventory.ProductList{}
	for i := range prodList.Products {
		if prodList.Products[i].ProductID == productID {
			for j := range prodList.Products[i].Variants {
				if prodList.Products[i].Variants[j].VariantID != variantID {
					currProdList.Products[i].Variants = append(currProdList.Products[i].Variants, prodList.Products[i].Variants[j])
				}
			}
		}
	}

	err = WriteToProducts(currProdList)
	if err != nil {
		return err
	}

	return nil

}