package routes

import (
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	"GOIOS/docs"
	"GOIOS/src/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() {

	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Akshit - Video API"
	docs.SwaggerInfo.Description = "Rest API in golang following best practices, built with gin, gorm(sqlite), swagger and MVC architecture."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"

	route := gin.New()

	apiRoutes := route.Group(docs.SwaggerInfo.BasePath)
	{
		v1 := apiRoutes.Group("/api/v1")
		{
			v1.POST("/user", controllers.CreateUser)
			v1.GET("/user", controllers.GetUser)
		}
	}
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	route.Run()
}
