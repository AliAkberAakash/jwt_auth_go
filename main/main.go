package main

import (
	"jwt-auth/main/src/controller"
	"jwt-auth/main/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var loginService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	server := gin.New()

	server.POST("/login", loginHandler(loginController))
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
