package models

type CartItems struct {
	ID        int64 `db:"id"`
	CartId    int64 `db:"cart_id"`
	ProductId int64 `db:"product_id"`
	Quantity  int   `db:"quantity"`
}
