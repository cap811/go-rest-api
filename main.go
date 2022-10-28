package main

import (
	"github.com/cap811/go-rest-api/initializers"
	"github.com/cap811/go-rest-api/routers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()
	routers.Fille()
	err := r.Run()
	if err != nil {
		return
	}
}
