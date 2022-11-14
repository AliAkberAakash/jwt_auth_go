package controller

import "jwt-auth/src/dto"

type UserController interface {
	GetAllUsers() []dto.User
}
