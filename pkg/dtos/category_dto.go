package dtos

type CategoryDTO struct {
	ID          int64        `json:"id,omitempty"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Products    []ProductDTO `json:"products,omitempty"`
}

type AllCategoryDTO struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}
