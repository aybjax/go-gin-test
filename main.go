package main

import (
	"github.com/aybjax/go-gin-test/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", handlers.Pong)

	router.Run(":8000")
}
