package products

type ProductEntity struct {
	ID                    string  `json:"id"`
	ProductName           string  `json:"product_name"`
	ProductCategory       string  `json:"product_category"`
	ProductFreshness      string  `json:"product_freshness"`
	ImageOfProduct        string  `json:"image_of_product"`
	AdditionalDescription string  `json:"additional_description"`
	ProductPrice          float64 `json:"product_price"`
}
