package models

import "time"

type Order struct {
	ID         int64     `db:"id"`
	UserId     int64     `db:"user_id"`
	OrderDate  time.Time `db:"order_date"`
	Status     string    `db:"status"` // pending, shipped, completed, cancelled, declined
	TotalPrice int64     `db:"total_price"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
