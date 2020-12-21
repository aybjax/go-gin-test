package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Pong is health test /ping handler
func Pong(c *gin.Context) {
	fmt.Printf("%#v\n\tis context", c)
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
