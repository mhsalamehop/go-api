package models

type UserModel struct {
	Id       string `json:"id"`
	Email    string `josn:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `jsong:"token"`
}
