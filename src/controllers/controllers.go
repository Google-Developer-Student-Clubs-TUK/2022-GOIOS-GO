package controllers

import (
	"GOIOS/src/config"
	"GOIOS/src/models"
	"crypto/rand"
	"errors"
	"io"
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

type CreateGroupRequest struct {
	UserName  string
	GroupName string
}

type JoinGroupRequest struct {
	UserName  string
	GroupCode string
}

// Create todo data to database by run this function
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

func CreateGroup(context *gin.Context) {
	var userData models.User
	var groupData models.Group
	var userGroupData models.UserGroup
	var requestData CreateGroupRequest

	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.Where("user_name = ?", requestData.UserName).First(&userData).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, gin.H{"result": "Fail", "message": "일치하는 회원 정보가 없습니다."})
			return
		}
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error getting data"})
		return
	}

	groupData.Name = requestData.GroupName
	groupData.GroupCode = EncodeToString(6)

	result := db.Create(&groupData)

	result.First(&groupData)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	userGroupData.GroupID = groupData.ID
	userGroupData.UserID = userData.ID
	userGroupData.Master = userData.UserName

	joinResult := db.Create(&userGroupData)

	if joinResult.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message":   "Success",
		"groupName": groupData.Name,
		"groupCode": groupData.GroupCode,
		"groupID":   groupData.ID,
	})

}

func JoinGroup(context *gin.Context) {
	var userData models.User
	var groupData models.Group
	var userGroupData models.UserGroup
	var requestData JoinGroupRequest

	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.Take(&groupData, "Group Code = ?", requestData.GroupCode).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, gin.H{"result": "Fail", "message": "그룹정보를 불러올 수  없습니다."})
			return
		}
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error getting data"})
		return
	}

	userGroupData.GroupID = groupData.ID
	userGroupData.UserID = userData.ID

	joinResult := db.Create(&userGroupData)

	if joinResult.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message":   "Join Success",
		"groupName": groupData.Name,
		"groupCode": groupData.GroupCode,
		"groupID":   groupData.ID,
	})

}

func EncodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
