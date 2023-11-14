package dto

type RegisterResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      uint   `json:"age"`
}
