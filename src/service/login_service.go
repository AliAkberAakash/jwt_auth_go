package service

import (
	"jwt-auth/src/dto"
	"jwt-auth/src/util"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginService interface {
	IsUserValid(loginRequest dto.LoginRequest) (bool, error)
}

type loginService struct {
	DB *gorm.DB
}

func StaticLoginService(db *gorm.DB) LoginService {
	return &loginService{
		DB: db,
	}
}

func (service *loginService) IsUserValid(loginRequest dto.LoginRequest) (bool, error) {

	foundUser, err := util.GetUserFromDB(loginRequest.Email, service.DB)

	if err != nil {
		return false, err
	}

	passErr := bcrypt.CompareHashAndPassword(
		[]byte(foundUser.Password),
		[]byte(loginRequest.Password),
	)
	if passErr != nil {
		log.Println(passErr)
		return false, passErr
	}
	return true, nil
}
