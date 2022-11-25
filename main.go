package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang_app/config"
	"golang_app/controller"
	docs "golang_app/docs"
	"golang_app/middleware"
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

// @title           Gin Book Service
// @version         1.0

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	api := r.Group("api")
	{
		authRoute := api.Group("auth")
		{
			authRoute.POST("/login", authController.Login)
			authRoute.POST("/register", authController.Register)
		}
		userRoute := api.Group("user").Use(middleware.AuthoizeJWT(jwtService))
		{
			userRoute.GET("/profile", userController.Profile)
		}
	}
	err := r.Run(":8081")
	if err != nil {
		return
	}
}
