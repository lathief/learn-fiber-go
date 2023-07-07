package models

type OrderProduct struct {
	ID       int64     `db:"id"`
	Products []Product `db:"products"`
	Order    Order     `db:"orders"`
}
