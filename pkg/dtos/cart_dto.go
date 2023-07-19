package dtos

type CartProductIdDTO struct {
	ProductId int `json:"product_id,omitempty"`
}
type CartDTO struct {
	ID     int64  `json:"id,omitempty"`
	Items  []Item `json:"items,omitempty"`
	UserId int64  `json:"user_id,omitempty"`
}
type CartUserDTO struct {
	ID        int64   `db:"id"`
	UserId    int64   `db:"user_id"`
	CartItems CartDTO `db:"cart_items"`
}
