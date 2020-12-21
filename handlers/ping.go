package handlers

import "github.com/gin-gonic/gin"

// Pong is health test /ping handler
func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
