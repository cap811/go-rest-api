package main

import (
	"github.com/cap811/go-rest-api/initializers"
	"github.com/cap811/go-rest-api/routers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	r := routers.Routes()
	err := r.Run()
	if err != nil {
		return
	}
}
