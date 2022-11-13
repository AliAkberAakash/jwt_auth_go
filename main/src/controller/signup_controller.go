package controller

import (
	"fmt"
	"jwt-auth/main/src/dto"
	"jwt-auth/main/src/service"

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
		return false, fmt.Errorf("No Data found")
	}

	return sc.signupService.Signup(user), nil
}
