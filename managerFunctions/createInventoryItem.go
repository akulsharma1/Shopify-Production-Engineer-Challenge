package managerfunctions

import (
	"encoding/json"
	"io/ioutil"
	"shopifyproductionengineertest/inventory"
)
func CreateInventoryItem(productName string, variants []inventory.Variant) bool {
	worked, currentProdList := ReadProducts()
	if worked {
		
		var newProd = inventory.Product{
			ProductName: productName,
			Variants:    variants,
		}
		
		currentProdList.Products = append(currentProdList.Products, newProd)
		file, _ := json.MarshalIndent(currentProdList, "", "	")
		err := ioutil.WriteFile("./inventory/inventory.json", file, 0644)
		return err==nil
	} else {
		return false
	}
	
}