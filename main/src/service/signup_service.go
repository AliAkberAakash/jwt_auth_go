package service

import "jwt-auth/main/src/dto"

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
		Users = append(Users, user)
		return true
	}

	return false
}
