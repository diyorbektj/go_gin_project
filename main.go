package main

import (
	"github.com/gin-gonic/gin"
	"golang_app/config"
	"golang_app/controller"
	"golang_app/repository"
	"golang_app/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	userService    service.UserService       = service.NewUserService(userRepository)
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	api := r.Group("api")
	{
		authRoute := api.Group("auth")
		{
			authRoute.POST("/login", authController.Login)
			authRoute.POST("/register", authController.Register)
		}
		userRoute := api.Group("user")
		{
			userRoute.GET("/profile", userController.Profile)
		}
	}
	err := r.Run()
	if err != nil {
		return
	}
}
