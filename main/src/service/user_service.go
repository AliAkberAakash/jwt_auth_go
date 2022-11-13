package service

import (
	"jwt-auth/main/src/data"
	"jwt-auth/main/src/dto"
)

type UserService interface {
	GetAllUser() []dto.User
}

type userService struct {
	users []dto.User
}

func NewUserService() UserService {
	return &userService{
		users: data.Users,
	}
}

func (service *userService) GetAllUser() []dto.User {
	return service.users
}
