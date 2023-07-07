package models

import "time"

type Order struct {
	ID        int64     `db:"id"`
	ProductId int64     `db:"product_id"`
	UserId    int64     `db:"user_id"`
	OrderDate time.Time `db:"order_date"`
	Status    string    `db:"status"` // pending, shipped, completed, cancelled, declined
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}