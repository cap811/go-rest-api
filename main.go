package main

import (
	"github.com/cap811/go-rest-api/initializers"
	"github.com/cap811/go-rest-api/routers"
	"github.com/gin-contrib/cors"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	r := routers.Routes()
	r.Use(cors.Default())
	err := r.Run("localhost:80")
	if err != nil {
		return
	}
}
