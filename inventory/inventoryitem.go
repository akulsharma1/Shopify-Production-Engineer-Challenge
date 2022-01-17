package inventory

type Product struct {
	ProductName string `bson:"productName" json:"productName"`
	ProductID   string `json:"productID" bson:"productID"`
	Variants1   []struct {
		VariantID   string  `json:"variantID" bson:"variantID"`
		VariantName string  `json:"variantName" bson:"variantName"`
		Size        string  `json:"size" bson:"size"`
		Image       string  `json:"image" bson:"image"`
		Quantity    int     `json:"quantity" bson:"quantity"`
		Price       float64 `json:"price" bson:"price"`
	} `json:"variants" bson:"variants"`
	Variants []Variant `json:"variants2" bson:"variants2"`
}
type Variant struct {
	VariantID   string  `json:"variantID" bson:"variantID"`
	VariantName string  `json:"variantName" bson:"variantName"`
	Size        string  `json:"size" bson:"size"`
	Image       string  `json:"image" bson:"image"`
	Quantity    int     `json:"quantity" bson:"quantity"`
	Price       float64 `json:"price" bson:"price"`
}
type ProductList struct {
	Products []Product `json:"products" bson:"products"`
}
