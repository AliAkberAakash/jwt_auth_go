package controller

import (
	"fmt"
	"jwt-auth/src/data"
	"jwt-auth/src/dto"
	"jwt-auth/src/service"
	"jwt-auth/src/util"

	"github.com/gin-gonic/gin"
)

type SignupController interface {
	Signup(ctx *gin.Context) (bool, error)
}

type signupController struct {
	signupService service.SignupService
}

func NewSignupController(signupService service.SignupService) SignupController {
	return &signupController{
		signupService: signupService,
	}
}

func (sc *signupController) Signup(ctx *gin.Context) (bool, error) {

	var user dto.User

	err := ctx.ShouldBind(&user)
	if err != nil {
		fmt.Println(err)
		return false, fmt.Errorf("Invalid or incorrect data")
	}

	valid := len(user.Email) > 0 && len(user.Password) > 8

	if !valid {
		return false, fmt.Errorf("Invalid or incorrect data")
	}

	if len(data.Users) != 0 {
		foundUser, err := util.GetUserFromDB(user.Email)

		if err == nil {
			return false,
				fmt.Errorf("User with email %s already exists", foundUser.Email)
		}
	}

	return sc.signupService.Signup(user), nil
}
