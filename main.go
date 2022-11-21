package main

import (
	"jwt-auth/src/controller"
	"jwt-auth/src/data"
	"jwt-auth/src/handler"
	"jwt-auth/src/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := data.InitStore()
	if err != nil {
		log.Fatalf("failed to initialise the store: %s", err)
	}

	var jwtService service.JWTService = service.JWTAuthService()

	var loginService service.LoginService = service.StaticLoginService(db)
	var loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)

	var signupService service.SignupService = service.StaticSignupService(
		db,
	)
	var signupController controller.SignupController = controller.NewSignupController(signupService)

	server := gin.New()

	server.POST("/login", handler.LoginHandler(loginController))
	server.POST("/signup", handler.SignUpHandler(signupController))
	//server.GET("/user", userHandler(signupController))

	port := "8080"
	server.Run(":" + port)

}
