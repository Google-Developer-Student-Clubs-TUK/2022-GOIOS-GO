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

// Define database client
var db *gorm.DB = config.ConnectDB()
var userdata []models.User

// Defining struct for response
type UserResponse struct {
	models.User
}

// Welcome godoc
// @Summary Summary를 적어 줍니다.
// @Description 자세한 설명은 이곳에 적습니다.
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Param name path string true "User name"
// @Router /user [post]
// @Success 200 {object} welcomeModel

// Create todo data to database by run this function
func CreateUser(context *gin.Context) {
	var data models.User

	// Binding request body json to request body struct
	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Matching todo models struct with todo request struct
	user := models.User{}
	user.UserName = data.UserName
	user.Password = data.Password

	// Querying to database
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

	// Matching result to create response
	var response UserResponse
	response.UserName = user.UserName
	response.Password = user.Password

	// Creating http response
	context.JSON(http.StatusCreated, response)
}

// Welcome godoc
// @Summary Summary를 적어 줍니다.
// @Description 자세한 설명은 이곳에 적습니다.
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Param name path string true "User name"
// @Router /user [get]
// @Success 200 {object} welcomeModel

func GetUser(context *gin.Context) {

	// Querying to find todo datas.
	err := db.Find(&userdata)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error getting data"})
		return
	}

	// Creating http response
	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    userdata,
	})

}
