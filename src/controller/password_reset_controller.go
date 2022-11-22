package controller

import (
	"fmt"
	"jwt-auth/src/dto"
	"jwt-auth/src/service"

	"github.com/gin-gonic/gin"
)

type PasswordResetController interface {
	SendPasswordResetCode(ctx *gin.Context) error
}

type passwordResetController struct {
	Service service.PasswordResetService
}

func GetPasswordResetController(service service.PasswordResetService) PasswordResetController {
	return &passwordResetController{
		Service: service,
	}
}

func (controller *passwordResetController) SendPasswordResetCode(ctx *gin.Context) error {
	var request dto.PasswordResetRequst

	err := ctx.ShouldBind(&request)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("Invalid or incorrect data")
	}

	return controller.Service.SendPasswordResetCode(request.Email)
}
