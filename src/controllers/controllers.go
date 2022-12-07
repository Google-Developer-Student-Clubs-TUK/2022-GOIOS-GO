package controllers

import (
	"GOIOS/src/config"
	"GOIOS/src/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)


var db *gorm.DB = config.ConnectDB()
var userdata []models.User


type UserResponse struct {
	models.User
}

// Paths Information

// Authenticate godoc
// @Summary Create users
// @Description Create new users
// @name get-string-by-int
// @Consume application/x-www-form-urlencoded
// @Produce json
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Router /api/v1/user [post]
// @Success 200 {object} models.Response
// @Failure 401 {object} models.Response
func CreateUser(context *gin.Context) {
	var data models.User

	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}
	user.UserName = data.UserName
	user.Password = data.Password

	result := db.Create(&user)
	if result.Error != nil {
		if pgError := result.Error.(*pgconn.PgError); errors.Is(result.Error, pgError) {
			switch pgError.Code {
			case "23505":
				context.JSON(http.StatusConflict, gin.H{"error": "same data"})
				return
			}
		}
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	var response UserResponse
	response.UserName = user.UserName
	response.Password = user.Password

	context.JSON(http.StatusCreated, response)
}

// GetUser godoc
// @Summary Get users
// @Description Get auth users
// @name get-string-by-int
// @Consume application/x-www-form-urlencoded
// @Produce json
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Router /api/v1/user [get]
// @Success 200 {object} models.Response
// @Failure 401 {object} models.Response
func GetUser(context *gin.Context) {

	err := db.Find(&userdata)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error getting data"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    userdata,
	})

}
