package models

import "time"

type Role struct {
	ID          int64     `db:"id"`
	RoleName    string    `db:"roleName"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
