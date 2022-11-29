package main

import (
	"go_app/config"
	"go_app/router"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
)

// @title           Gin Book Service
// @version         1.0

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	defer config.CloseDatabaseConnection(db)
	r := router.SetupRouter()
	err := r.Run(":5500")
	if err != nil {
		return
	}
}
