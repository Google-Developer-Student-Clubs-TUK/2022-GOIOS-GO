package routes

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"GOIOS/docs"
	"GOIOS/src/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// swagger API 선언
func setupSwagger(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/swagger/index.html")
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func Routes() {
	route := gin.Default()
	setupSwagger(route)
	v1 := route.Group("/api/v1")
	{
		docs.SwaggerInfo.Title = "Swagger Example API"
		docs.SwaggerInfo.Description = "This is a sample server for Swagger."
		docs.SwaggerInfo.Version = "1.0"
		docs.SwaggerInfo.Host = "localhost:8080"
		docs.SwaggerInfo.BasePath = "/api/v1"

		v1.POST("/user", controllers.CreateUser)
		v1.GET("/user", controllers.GetUser)
		// route.GET("/todo", controllers.GetAllTodos)
		// route.PUT("/todo/:idTodo", controllers.UpdateTodo)
		// route.DELETE("/todo/:idTodo", controllers.DeleteTodo)
	}

	route.Run()
}
