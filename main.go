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
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	authRoute := r.Group("api/auth")
	{
		authRoute.POST("/login", authController.Login)
		authRoute.POST("/register", func(ctx *gin.Context) {})
	}
	r.Run()
}
