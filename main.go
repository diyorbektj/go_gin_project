package main

import (
  "golang_app/config"
  "github.com/gin-gonic/gin"
  "gorm.io/gorm"
  "golang_app/repository"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
  userRepository repository.UserRepository = repository.NewUserRepository(db)
)


func main() {
  defer config.CloseDatabaseConnection(db)
  r := gin.Default()
  authRoute := r.Group("api/auth")
	{
		authRoute.POST("/login", func(ctx *gin.Context) {})
		authRoute.POST("/register", func(ctx *gin.Context) {})
	}
  r.Run()
}