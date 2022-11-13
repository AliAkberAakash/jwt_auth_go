package service

import (
	"fmt"
	"jwt-auth/main/src/dto"
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

	foundUser, err := getUserFromDB(loginRequest.Email)

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

func getUserFromDB(email string) (dto.User, error) {

	var foundUser dto.User

	for _, u := range Users {
		if u.Email == email {
			foundUser = u
			return foundUser, nil
		}
	}

	return foundUser, fmt.Errorf("User not Found")
}
