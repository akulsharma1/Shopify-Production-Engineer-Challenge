package inventory

type Product struct {
	ProductName string    `json:"productName"`
	ProductID   string    `json:"productID"`
	Variants    []Variant `json:"variants"`
}
type Variant struct {
	VariantID   int     `json:"variantID"`
	VariantName string  `json:"variantName"`
	Size        string  `json:"size"`
	Image       string  `json:"image"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}
type ProductList struct {
	Products []Product `json:"products"`
}