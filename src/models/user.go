package models

type User struct {
	Id       uint   `json:"id"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password []byte `json:"-"`
	RoleId   uint   `json:"role_id"`
	Role     Role   `json:"role"`
}
