package service

import "jwt-auth/main/src/dto"

type UserService interface {
	GetAllUser() []dto.User
}

type userService struct {
	users []dto.User
}

func NewUserService() UserService {
	return &userService{
		users: Users,
	}
}

func (service *userService) GetAllUser() []dto.User {
	return service.users
}
