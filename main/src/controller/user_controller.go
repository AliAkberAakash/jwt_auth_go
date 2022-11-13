package controller

import "jwt-auth/main/src/dto"

type UserController interface {
	GetAllUsers() []dto.User
}
