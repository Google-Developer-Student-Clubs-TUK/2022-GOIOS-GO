package routes

import (
	"GOIOS/src/controllers"

	"github.com/gin-gonic/gin"
)

// Routes function to serve endpoints
func Routes() {
	route := gin.Default()

	v1 := route.Group("/api/v1")
	{
		v1.POST("/user", controllers.CreateUser)
		v1.GET("/user", controllers.GetUser)
		v1.POST("/user/login", controllers.Login)
		// route.GET("/todo", controllers.GetAllTodos)
		// route.PUT("/todo/:idTodo", controllers.UpdateTodo)
		// route.DELETE("/todo/:idTodo", controllers.DeleteTodo)
	}
	// Run route whenever triggered
	route.Run()
}
