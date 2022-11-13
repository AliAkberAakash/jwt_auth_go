package service

import (
	"jwt-auth/main/src/dto"
	"jwt-auth/main/src/util"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type LoginService interface {
	IsUserValid(loginRequest dto.LoginRequest) bool
}

type loginService struct{}

func StaticLoginService() LoginService {
	return &loginService{}
}

func (info *loginService) IsUserValid(loginRequest dto.LoginRequest) bool {

	foundUser, err := util.GetUserFromDB(loginRequest.Email)

	if err != nil {
		log.Println(err)
		return false
	}

	passErr := bcrypt.CompareHashAndPassword(
		[]byte(foundUser.Password),
		[]byte(loginRequest.Password),
	)
	if passErr != nil {
		log.Println(passErr)
		return false
	}
	return true
}
