package routes

import (
	"webapi-with-go/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1") //Seta como path após a url principal (localhost:5000/api/v1)
	{
		books := main.Group("books") //localhost:5000/api/v1/books
		{
			books.GET("/", controllers.ShowBook)
		}
	}
	return router
}
