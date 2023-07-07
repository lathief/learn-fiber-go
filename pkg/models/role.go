package models

type Role struct {
	ID          int64  `json:"id"`
	RoleName    string `json:"roleName"`
	Description string `json:"description"`
}
