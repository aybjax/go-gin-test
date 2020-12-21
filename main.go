package main

import (
	"github.com/aybjax/go-gin-test/database"
	"github.com/aybjax/go-gin-test/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	defer database.Close()

	router := gin.Default()

	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// 	v.RegisterStructValidation()
	// }

	router.GET("/ping", handlers.Pong)
	router.POST("/save", handlers.Save)

	router.Run(":8000")
}
