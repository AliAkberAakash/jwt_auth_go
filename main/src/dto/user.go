package dto

type User struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	DOB      string `json:"dateOfBirth,omitempty"`
	Address  string `json:"address,omitempty"`
}
