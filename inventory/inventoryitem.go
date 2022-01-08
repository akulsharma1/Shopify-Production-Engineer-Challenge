package inventory

type Product struct {
	ProductName string `json:"productName"`
	Sizes       []struct {
		SizeID   int    `json:"sizeID"`
		Size     string `json:"size"`
		Name     string `json:"name"`
		Image    string `json:"image"`
		Quantity int    `json:"quantity"`
	} `json:"sizes"`
}

type ProductList struct {
	Products []Product `json:"products"`
}