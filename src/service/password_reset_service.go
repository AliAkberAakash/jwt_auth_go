package service

import (
	"jwt-auth/src/dto"
	"jwt-auth/src/util"
	"log"
	"time"

	"gorm.io/gorm"
)

type PasswordResetService interface {
	SendPasswordResetCode(email string) error
}

type passwordResetService struct {
	DB *gorm.DB
}

func GetPasswordResetRequestService(db *gorm.DB) PasswordResetService {
	return &passwordResetService{
		DB: db,
	}
}

func (service *passwordResetService) SendPasswordResetCode(email string) error {

	foundUser, err := util.GetUserFromDB(email, service.DB)
	if err != nil {
		return err
	}

	token := generateToken()

	err = saveTokenInDatabase(token, *foundUser, service.DB)
	if err != nil {
		return err
	}

	err = sendTokenToEmail(token, email)
	if err != nil {
		return err
	}

	return nil
}

func saveTokenInDatabase(token string, user dto.User, db *gorm.DB) error {
	validity := time.Now().Add(time.Minute * 15)

	// check if entry for user already exists
	var existingData dto.PasswordResetToken
	result := db.Where("user_id = ?", user.ID).First(&existingData)
	var count int64
	result.Count(&count)

	// if no previous entry found
	// insert new
	// else update previous
	if count == 0 {
		resetTokenData := &dto.PasswordResetToken{
			UserID:    user.ID,
			Token:     token,
			ExpiresAt: validity,
		}
		db.Create(&resetTokenData)
	} else {
		existingData.Token = token
		existingData.ExpiresAt = validity
		db.Save(&existingData)
	}

	return nil
}

func sendTokenToEmail(token string, email string) error {
	log.Print("Sending token to email")
	return nil
}
