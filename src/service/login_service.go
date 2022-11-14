package service

import (
	"jwt-auth/src/dto"
	"jwt-auth/src/util"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type LoginService interface {
	IsUserValid(loginRequest dto.LoginRequest) (bool, error)
}

type loginService struct{}

func StaticLoginService() LoginService {
	return &loginService{}
}

func (info *loginService) IsUserValid(loginRequest dto.LoginRequest) (bool, error) {

	foundUser, err := util.GetUserFromDB(loginRequest.Email)

	if err != nil {
		log.Println(err)
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
