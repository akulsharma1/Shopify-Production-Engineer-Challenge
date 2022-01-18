package inventory

type Product struct {
	ProductName string    `bson:"productName" json:"productName"`
	ProductID   string    `json:"productID" bson:"productID"`
	Variants    []Variant `json:"variants" bson:"variants"`
}
type Variant struct {
	VariantID   string  `json:"variantID" bson:"variantID"`
	VariantName string  `json:"variantName" bson:"variantName"`
	Size        string  `json:"size" bson:"size"`
	Quantity    int     `json:"quantity" bson:"quantity"`
	Price       float64 `json:"price" bson:"price"`
}
type ProductList struct {
	Products []Product `json:"products" bson:"products"`
}

// fiber stuff
//
//
//

type FiberAddVariant struct {
	ProductID   string  `json:"productID"`
	VariantID   string  `json:"variantID" bson:"variantID"`
	VariantName string  `json:"variantName" bson:"variantName"`
	Size        string  `json:"size" bson:"size"`
	Image       string  `json:"image" bson:"image"`
	Quantity    int     `json:"quantity" bson:"quantity"`
	Price       float64 `json:"price" bson:"price"`
}

type FiberGetProductID struct {
	ProductID string `json:"productID"`
}

type FiberGetVariantInfo struct {
	ProductID string `json:"productID"`
	VariantID string `json:"variantID"`
}