package managerfunctions

import (
	"encoding/json"
	"io/ioutil"
	"shopifyproductionengineertest/inventory"
)

func WriteToProducts(productList inventory.ProductList) error {
	file, _ := json.MarshalIndent(productList, "", "	")
	err := ioutil.WriteFile("./inventory/inventory.json", file, 0644)
	if err != nil {
		return err
	}

	return nil
}