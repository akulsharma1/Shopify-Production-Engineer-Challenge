package managerfunctions

import (
	"shopifyproductionengineertest/inventory"
)
func CreateInventoryItem(productName string, variants []inventory.Variant) error {
	currentProdList, err := ReadProducts()
	if err != nil {
		return err
	} 
	var newProd = inventory.Product{
		ProductName: productName,
		Variants:    variants,
	}
	
	currentProdList.Products = append(currentProdList.Products, newProd)
	
	err = WriteToProducts(currentProdList)
	if err != nil {
		return err
	}

	return nil
}

func CreateInventoryVariant(productID, variantID, name, size, image string, quantity int, price float64) error {
	var variant = inventory.Variant{
		VariantID:   variantID,
		VariantName: name,
		Size:        size,
		Image:       image,
		Quantity:    quantity,
		Price:       price,
	}

	prodList, err := ReadProducts()
	if err != nil {
		return err
	}

	for i := range prodList.Products {
		if prodList.Products[i].ProductID == productID {
			prodList.Products[i].Variants = append(prodList.Products[i].Variants, variant)
		}
	}

	err = WriteToProducts(prodList)
	if err != nil {
		return err
	}

	return nil
}