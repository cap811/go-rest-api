package routers

import (
	"github.com/cap811/go-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.Default()

	// Routes
	api := r.Group("/api")
	{
		api.GET("/books", controllers.ShowAllBooks)
		api.GET("/books/:id", controllers.FindBook)
		api.POST("/books", controllers.CreateBooks)
		api.POST("/books/:id", controllers.UpdateBooks)
		api.DELETE("/books/:id", controllers.DeleteBooks)
	}
	return r
}
