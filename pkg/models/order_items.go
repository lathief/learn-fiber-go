package models

type OrderItems struct {
	ID        int64   `db:"id"`
	OrderId   int64   `db:"order_id"`
	ProductId int64   `db:"product_id"`
	Quantity  int     `db:"quantity"`
	Price     float64 `db:"price"`
}
