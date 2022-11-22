package service

import (
	"jwt-auth/src/util"

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

	_, err := util.GetUserFromDB(email, service.DB)

	if err != nil {
		return err
	}

	return nil
}
