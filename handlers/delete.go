package handlers

import (
	"net/http"

	"github.com/aybjax/go-gin-test/constants"
	"github.com/aybjax/go-gin-test/controller"
	"github.com/aybjax/go-gin-test/data_structures"
	"github.com/gin-gonic/gin"
)

var deleteData data_structures.DeleteData

// Delete -> delete token
func Delete(c *gin.Context) {
	deleteData = data_structures.DeleteData{}
	err := c.ShouldBindJSON(&deleteData)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.INCORRECT_DATA)
	}

	if !controller.DeleteData(&deleteData) {
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
