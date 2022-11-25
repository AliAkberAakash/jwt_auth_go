package service

import (
	"fmt"
	"jwt-auth/src/dto"
	"jwt-auth/src/util"

	"gorm.io/gorm"
)

type SignupService interface {
	Signup(user dto.User) error
}

type signupService struct {
	DB *gorm.DB
}

func StaticSignupService(db *gorm.DB) SignupService {
	return &signupService{
		DB: db,
	}
}

func (info *signupService) Signup(user dto.User) error {

	_, err := util.GetUserFromDB(user.Email, info.DB)

	if err == nil {
		return fmt.Errorf("User already exists with email %s", user.Email)
	}

	hashedPass := util.GetHash([]byte(user.Password))

	newUser := dto.User{
		Email:    user.Email,
		Password: hashedPass,
	}

	info.DB.Table("users").Create(&newUser)

	return nil
}
