package models

type UserModel struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token,omitempty"`
	Role     string `json:"role"`
}

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
