package dto

type PasswordResetRequst struct {
	Email string `json:"email" binding:"required"`
}
