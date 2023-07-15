package dtos

type CartDTO struct {
	ID     int64  `json:"id,omitempty"`
	Items  []Item `json:"items,omitempty"`
	UserId int64  `json:"user_id,omitempty"`
}
