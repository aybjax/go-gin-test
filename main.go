package main

import (
	"github.com/aybjax/go-gin-test/database"
	"github.com/aybjax/go-gin-test/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	defer database.Close()

	router := gin.Default()

	router.GET("/ping", handlers.Pong)
	router.POST("/", handlers.Save)
	router.DELETE("/", handlers.Delete)
	router.GET("/", handlers.Get)

	router.Run(":8000")
}
