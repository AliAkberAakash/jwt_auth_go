package dto

import (
	"time"

	"gorm.io/gorm"
)

type PasswordResetToken struct {
	gorm.Model
	UserID    uint
	Token     string
	ExpiresAt time.Time
}
