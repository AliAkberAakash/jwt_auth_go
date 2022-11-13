package service

import (
	"jwt-auth/main/src/dto"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type SignupService interface {
	Signup(user dto.User) bool
}

type signupInformation struct {
}

func StaticSignupService() SignupService {
	return &signupInformation{}
}

func (info *signupInformation) Signup(user dto.User) bool {
	valid := len(user.Email) > 0 && len(user.Password) > 8

	if valid {
		hashedPass := getHash([]byte(user.Password))

		newUser := dto.User{
			Email:    user.Email,
			Password: hashedPass,
		}

		Users = append(Users, newUser)
		return true
	}

	return false
}

func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
