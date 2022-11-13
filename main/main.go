package main

import (
	"jwt-auth/main/src/controller"
	"jwt-auth/main/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	var jwtService service.JWTService = service.JWTAuthService()

	var loginService service.LoginService = service.StaticLoginService()
	var loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)

	var signupService service.SignupService = service.StaticSignupService()
	var signupController controller.SignupController = controller.NewSignupController(signupService)

	server := gin.New()

	server.POST("/login", loginHandler(loginController))
	server.POST("/signup", signUpHandler(signupController))
	//server.GET("/user", userHandler(signupController))

	port := "8080"
	server.Run(":" + port)

}

func loginHandler(loginController controller.LoginController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	}
}

func signUpHandler(signupController controller.SignupController) gin.HandlerFunc {
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
