package main

import (
	"go_app/config"
	"gorm.io/gorm"
	"go_app/router"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
)

// @title           Gin Book Service
// @version         1.0

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	defer config.CloseDatabaseConnection(db)
	r := router.setupRouter()
	err := r.Run(":5500")
	if err != nil {
		return
	}
}