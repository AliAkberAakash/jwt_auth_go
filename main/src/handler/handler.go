package handler

import (
	"jwt-auth/main/src/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(loginController controller.LoginController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Success",
				"token":   token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Success",
			})
		}
	}
}

func SignUpHandler(signupController controller.SignupController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		isUserCreated, err := signupController.Signup(ctx)
		if err == nil && isUserCreated {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "User Created Successfully",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
	}
}
