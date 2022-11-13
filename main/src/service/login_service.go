package service

import (
	"fmt"
	"jwt-auth/main/src/dto"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type LoginService interface {
	IsUserValid(user dto.User) bool
}

type loginService struct{}

func StaticLoginService() LoginService {
	return &loginService{}
}

func (info *loginService) IsUserValid(user dto.User) bool {

	currentUser := dto.User{
		Email:    user.Email,
		Password: user.Password,
	}

	foundUser, err := getUserFromDB(currentUser)

	if err != nil {
		log.Println(err)
		return false
	}

	passErr := bcrypt.CompareHashAndPassword(
		[]byte(foundUser.Password),
		[]byte(user.Password),
	)
	if passErr != nil {
		log.Println(passErr)
		return false
	}
	return true
}

func getUserFromDB(currentUser dto.User) (dto.User, error) {

	var foundUser dto.User

	for _, u := range Users {
		if u.Email == currentUser.Email {
			foundUser = u
			return foundUser, nil
		}
	}

	return foundUser, fmt.Errorf("User not Found")
}
