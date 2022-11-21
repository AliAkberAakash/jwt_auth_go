package controller

import (
	"fmt"
	"jwt-auth/src/dto"
	"jwt-auth/src/service"

	"github.com/gin-gonic/gin"
)

type SignupController interface {
	Signup(ctx *gin.Context) error
}

type signupController struct {
	signupService service.SignupService
}

func NewSignupController(signupService service.SignupService) SignupController {
	return &signupController{
		signupService: signupService,
	}
}

func (sc *signupController) Signup(ctx *gin.Context) error {

	var user dto.User

	err := ctx.ShouldBind(&user)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("Invalid or incorrect data")
	}

	valid := len(user.Email) > 0 && len(user.Password) > 8

	if !valid {
		return fmt.Errorf("Invalid or incorrect data")
	}

	err = sc.signupService.Signup(user)

	return err
}
