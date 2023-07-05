package models

import "time"

type Product struct {
	ID          int64     `db:"product_id"`
	Name        string    `db:"product_name"`
	Description string    `db:"product_description"`
	Price       float64   `db:"product_price"`
	CategoryId  int64     `db:"product_category_id"`
	Category    Category  `db:"product_category"`
	CreatedAt   time.Time `db:"product_created_at"`
	UpdatedAt   time.Time `db:"product_updated_at"`
}
