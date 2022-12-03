package routes

import (
	"GOIOS/src/controllers"

	"github.com/gin-gonic/gin"
)

// Routes function to serve endpoints
func Routes() {
	route := gin.Default()

	route.POST("/api/v1/user", controllers.CreateUser)
	route.GET("/api/v1/user", controllers.GetUser)
	// route.GET("/todo", controllers.GetAllTodos)
	// route.PUT("/todo/:idTodo", controllers.UpdateTodo)
	// route.DELETE("/todo/:idTodo", controllers.DeleteTodo)

	// Run route whenever triggered
	route.Run()
}
