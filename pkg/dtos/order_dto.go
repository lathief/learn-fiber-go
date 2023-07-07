package dtos

import "time"

type OrderDTO struct {
	ID        int64        `json:"id,omitempty"`
	Products  []ProductDTO `json:"products"`
	UserId    int64        `json:"user_id,omitempty"`
	Status    string       `json:"status"`
	OrderDate time.Time    `json:"orderDate"`
}
type AllOrderDTO struct {
	ID        int64        `json:"id,omitempty"`
	Products  []ProductDTO `json:"products"`
	User      UserDTO      `json:"user"`
	Status    string       `json:"status"`
	OrderDate time.Time    `json:"orderDate"`
}
