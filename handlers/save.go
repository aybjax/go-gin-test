package handlers

import (
	"net/http"

	"github.com/aybjax/go-gin-test/controller"
	"github.com/aybjax/go-gin-test/data_structures"
	"github.com/gin-gonic/gin"
)

var tokenData data_structures.InsertData

// Save is handler of saving
// 		authorized and unauthorized data
func Save(c *gin.Context) {
	// zero initialize: or user_id is kept
	tokenData = data_structures.InsertData{}
	err := c.ShouldBindJSON(&tokenData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "incorrect data",
		})

		return
	}

	if !controller.InsertData(tokenData) {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "database error",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "ok",
		},
	)
}
