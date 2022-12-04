package main

import (
	"GOIOS/src/config"
	"GOIOS/src/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func init() {
	db.AutoMigrate(&models.Group{}, &models.UserGroup{}, &models.User{})

}

func main() {
	defer config.DisconnectDB(db)
	r := setupRouter()
	r.Run(":8080")
}
