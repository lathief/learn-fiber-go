package models

type User struct {
	ID          int64  `db:"id"`
	FirstName   string `db:"first_name"`
	LastName    string `db:"last_name"`
	Email       string `db:"email"`
	Password    string `db:"password"`
	Address     string `db:"address"`
	PhoneNumber string `db:"phone_number"`
	RoleId      int64  `db:"role_id"`
}
