package controllers

import (
	"errors"
	"github.com/jackc/pgconn"
	"net/http"

	"GOIOS/src/config"
	"GOIOS/src/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Define database client
var db *gorm.DB = config.ConnectDB()
var userdata []models.User

// Defining struct for response
type UserResponse struct {
	models.User
}

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
				context.JSON(http.StatusConflict, gin.H{"error": "conflict"})
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

func Login(context *gin.Context) {

	var data models.User

	// Binding request body json to request body struct
	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := db.Where("user_name = ?", data.UserName).First(&userdata).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, gin.H{"result": "Fail", "message": "일치하는 회원 정보가 없습니다."})
			return
		}
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error getting data"})
		return
	}

	if userdata[0].Password == data.Password {
		context.JSON(http.StatusOK, gin.H{"result": "Success", "user": userdata[0]})
		return
	} else {
		context.JSON(http.StatusNotAcceptable, gin.H{"result": "Fail", "message": "일치하는 회원 정보가 없습니다."})
		return
	}

}
