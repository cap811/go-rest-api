package routers

import (
	"github.com/cap811/go-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func Fille() {
	r := gin.Default()

	// Routes
	r.GET("/api/books", controllers.ShowAllBooks)
	r.GET("/api/books/:id", controllers.FindBook)
	r.POST("/api/books/", controllers.CreateBooks)
	r.POST("/api/books/:id", controllers.UpdateBooks)
}
