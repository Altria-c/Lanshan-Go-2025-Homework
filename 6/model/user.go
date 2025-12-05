package model

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ChangeUser struct {
	User
	NewPassword string `json:"newpassword"`
}
