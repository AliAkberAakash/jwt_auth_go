package handler

import (
	"jwt-auth/src/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(loginController controller.LoginController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := loginController.Login(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "success",
				"token":   token,
			})
		}
	}
}

func SignUpHandler(signupController controller.SignupController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := signupController.Signup(ctx)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Password reset code sent to email succeessfully",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
	}
}

func PasswordResetRequestHandler(passwordResetController controller.PasswordResetController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := passwordResetController.SendPasswordResetCode(ctx)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}
	}
}
