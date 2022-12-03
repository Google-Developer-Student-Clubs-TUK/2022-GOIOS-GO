package main

import (
	"GOIOS/src/config"
	"GOIOS/src/models"
	"GOIOS/src/routes"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func init() {
	db.AutoMigrate(&models.Group{}, &models.UserGroup{}, &models.User{})
	if err := db.SetupJoinTable(&models.User{}, "Groups", &models.UserGroup{}); err != nil {
		println(err.Error())
		panic("Failed to setup join table")
	}
}

func main() {
	defer config.DisconnectDB(db)
	routes.Routes()
}
