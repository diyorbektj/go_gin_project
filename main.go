package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go_app/config"
	"go_app/controller"
	_ "go_app/controller"
	docs "go_app/docs"
	"go_app/middleware"
	"go_app/repository"
	"go_app/service"
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
	homeController controller.HomeController = controller.NewHomeController()
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
	r.Use(middleware.ApiCors())
	r.GET("/", homeController.Home)
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
	err := r.Run(":5500")
	if err != nil {
		return
	}
}
