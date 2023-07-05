package models

import "time"

type Category struct {
	ID          int64     `db:"category_id"`
	Name        string    `db:"category_name"`
	Description string    `db:"category_description"`
	CreatedAt   time.Time `db:"category_created_at"`
	UpdatedAt   time.Time `db:"category_updated_at"`
}
