package service

import (
	"fmt"
	"jwt-auth/src/dto"
	"jwt-auth/src/util"
	"log"
	"time"

	"gorm.io/gorm"
)

type ResetPasswordService interface {
	SendResetPasswordCode(email string) error
	ResetPassword(request dto.ResetPasswordRequest) error
}

type resetPasswordService struct {
	DB *gorm.DB
}

func GetPasswordResetRequestService(db *gorm.DB) ResetPasswordService {
	return &resetPasswordService{
		DB: db,
	}
}

func (service *resetPasswordService) SendResetPasswordCode(email string) error {

	foundUser, err := util.GetUserFromDB(email, service.DB)
	if err != nil {
		return err
	}

	token := generateToken()
	hashedToken := token //util.GetHash([]byte(token))

	err = saveTokenInDatabase(hashedToken, *foundUser, service.DB)
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
	log.Printf("Token: %s", token)
	return nil
}

func (service *resetPasswordService) ResetPassword(request dto.ResetPasswordRequest) error {

	hashedToken := request.Token //util.GetHash([]byte(request.Token))
	var existingData dto.PasswordResetToken
	result := service.DB.Table("password_reset_tokens").Where("token = ?", hashedToken).First(&existingData)

	var count int64
	result.Count(&count)

	if count == 0 {
		return fmt.Errorf("Invalid token")
	}

	hashedPassword := util.GetHash([]byte(request.Password))

	result = service.DB.Table("users").Where("ID = ?", existingData.UserID).Update("password", hashedPassword)

	if result.Error != nil {
		log.Print(result.Error)
		return fmt.Errorf("Failed to update password")
	}

	return nil
}
