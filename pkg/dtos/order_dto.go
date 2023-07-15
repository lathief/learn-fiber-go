package dtos

import "time"

type OrderDTO struct {
	ID         int64        `json:"id,omitempty"`
	Products   []ProductDTO `json:"products"`
	UserId     int64        `json:"user_id"`
	Status     string       `json:"status"`
	TotalPrice int64        `db:"total_price"`
	OrderDate  time.Time    `json:"orderDate"`
}
type AllOrderDTO struct {
	ID         int64     `json:"id,omitempty"`
	User       int64     `json:"user_id"`
	Status     string    `json:"status"`
	TotalPrice int64     `db:"total_price"`
	OrderDate  time.Time `json:"orderDate"`
}
type OrderReqDTO struct {
	ID         int64     `json:"id,omitempty"`
	ProductsId []int64   `json:"products_id"`
	UserId     int64     `json:"user_id"`
	Status     string    `json:"status"`
	OrderDate  time.Time `json:"orderDate"`
}
