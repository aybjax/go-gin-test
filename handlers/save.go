package handlers

import (
	"net/http"

	"github.com/aybjax/go-gin-test/constants"
	"github.com/aybjax/go-gin-test/controller"
	"github.com/aybjax/go-gin-test/data_structures"
	"github.com/gin-gonic/gin"
)

var insertData data_structures.InsertData

// Save is handler of saving
// 		authorized and unauthorized data
func Save(c *gin.Context) {
	// zero initialize: or user_id is kept
	insertData = data_structures.InsertData{}
	err := c.ShouldBindJSON(&insertData)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.INCORRECT_DATA)

		return
	}

	if !controller.InsertData(&insertData) {
		c.JSON(
			http.StatusInternalServerError,
			constants.DATABASE_ERROR,
		)
		return
	}

	c.JSON(
		http.StatusOK,
		constants.STATUS_OK,
	)
}
