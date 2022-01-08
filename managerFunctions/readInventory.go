package managerfunctions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"shopifyproductionengineertest/inventory"
)
func ReadProducts() bool {
	var prodListStruct inventory.ProductList
	jsonFile, err := os.Open("./inventory/inventory.json")

	if err != nil {
		fmt.Println(err)
		return false
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &prodListStruct)

	for i := range prodListStruct.Products {
		fmt.Println(prodListStruct.Products[i])
	}

	return true
}