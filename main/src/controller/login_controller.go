package controller

import (
	"fmt"
	"jwt-auth/main/src/dto"
	"jwt-auth/main/src/service"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	LoginService service.LoginService
	JwtService   service.JWTService
}

func NewLoginController(
	loginService service.LoginService,
	jwtService service.JWTService,
) LoginController {
	return &loginController{
		LoginService: loginService,
		JwtService:   jwtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var loginRequest dto.LoginRequest
	err := ctx.ShouldBind(&loginRequest)

	if err != nil {
		fmt.Println(err)
		return "no data found"
	}

	isUserAuthenticated := controller.LoginService.IsUserValid(loginRequest)
	if isUserAuthenticated {
		return controller.JwtService.GenerateToken(loginRequest.Email, true)
	}
	return ""
}
