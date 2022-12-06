package service

import (
	"fmt"
	"jwt-auth/src/dto"
	"jwt-auth/src/util"
	"log"
	"os"
	"time"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
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

	err = saveTokenInDatabase(hashedToken, foundUser.ID, service.DB)
	if err != nil {
		return err
	}

	err = sendTokenToEmail(token, *foundUser)
	if err != nil {
		return err
	}

	return nil
}

func saveTokenInDatabase(token string, uid uint, db *gorm.DB) error {
	validity := time.Now().Add(time.Minute * 15)

	// check if entry for user already exists
	var existingData dto.PasswordResetToken
	result := db.Where("user_id = ?", uid).First(&existingData)
	var count int64
	result.Count(&count)

	// if no previous entry found
	// insert new
	// else update previous
	if count == 0 {
		resetTokenData := &dto.PasswordResetToken{
			UserID:    uid,
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

func sendTokenToEmail(token string, user dto.User) error {
	from := mail.NewEmail("Hatil", "cyberwortsoftwares@gmail.com")
	subject := "Hatil Password Reset Token"
	to := mail.NewEmail(user.Email, user.Email)
	plainTextContent := "Your passwrd reset token is "
	htmlContent := "<strong>" + token + "</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	api_key := os.Getenv("SENDGRID_API_KEY")
	client := sendgrid.NewSendClient(api_key)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
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

	if existingData.ExpiresAt.Before(time.Now()) {
		return fmt.Errorf("Token has expired")
	}

	// replace the old token
	token := generateToken()
	err := saveTokenInDatabase(token, existingData.UserID, service.DB)
	if err != nil {
		return err
	}

	hashedPassword := util.GetHash([]byte(request.Password))
	result = service.DB.Table("users").Where("ID = ?", existingData.UserID).Update("password", hashedPassword)
	if result.Error != nil {
		log.Print(result.Error)
		return fmt.Errorf("Failed to update password")
	}

	return nil
}
