package controller

import (
	"fmt"
	"jwt-auth/src/dto"
	"jwt-auth/src/service"

	"github.com/gin-gonic/gin"
)

type ResetPasswordController interface {
	SendPasswordResetCode(ctx *gin.Context) error
	ResetPassword(ctx *gin.Context) error
}

type resetPasswordController struct {
	Service service.ResetPasswordService
}

func GetResetPasswordController(service service.ResetPasswordService) ResetPasswordController {
	return &resetPasswordController{
		Service: service,
	}
}

func (controller *resetPasswordController) SendPasswordResetCode(ctx *gin.Context) error {
	var request dto.ForgetPasswordRequest

	err := ctx.ShouldBind(&request)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("Invalid or incorrect data")
	}

	return controller.Service.SendResetPasswordCode(request.Email)
}

func (controller *resetPasswordController) ResetPassword(ctx *gin.Context) error {
	var request dto.ResetPasswordRequest

	err := ctx.ShouldBind(&request)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("Invalid or incorrect data")
	}

	if request.Password != request.ConfirmPassword {
		return fmt.Errorf("Passwords do not match")
	}

	if len(request.Password) < 8 {
		return fmt.Errorf("Passwords length too short")
	}

	return controller.Service.ResetPassword(request)
}
