package dtos

type UserDTO struct {
	ID          int64  `json:"id,omitempty"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	RoleId      int64  `json:"role_id,omitempty"`
}
