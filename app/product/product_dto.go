package product

type ProductDTO struct {
	ID          int64   `json:"id,omitempty"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryId  int64   `json:"category_id,omitempty"`
	Category    any     `json:"category,omitempty"`
}

type AllProductsDTO struct {
	ID          int64   `json:"id,omitempty"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryId  int64   `json:"category_id,omitempty"`
}
