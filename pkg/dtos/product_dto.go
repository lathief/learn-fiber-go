package dtos

type ProductDTO struct {
	ID          int64       `json:"id,omitempty"`
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty"`
	Price       float64     `json:"price"`
	CategoryId  int64       `json:"category_id,omitempty"`
	Category    CategoryDTO `json:"category,omitempty"`
}

type AllProductsDTO struct {
	ID          int64   `json:"id,omitempty"`
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price"`
	CategoryId  int64   `json:"category_id,omitempty"`
}

type Item struct {
	ProductId int64 `json:"product_id"`
	Quantity  int   `json:"quantity"`
}
