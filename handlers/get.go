package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aybjax/go-gin-test/constants"
	"github.com/aybjax/go-gin-test/controller"
	"github.com/aybjax/go-gin-test/data_structures"
	"github.com/gin-gonic/gin"
)

var getData data_structures.GetData

// Get -> get token
func Get(c *gin.Context) {
	getData = data_structures.GetData{}
	err := c.ShouldBindJSON(&getData)

	if err != nil {
		c.JSON(http.StatusBadRequest, constants.INCORRECT_DATA)
		fmt.Println("\n\n", err.Error(), "\n\n")
		return
	}

	tokens := controller.GetData(&getData)

	if tokens == nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "ok",
				"data":    nil,
			},
		)
		return
	}

	jsonTokens, err := json.Marshal(tokens)

	c.JSON(
		http.StatusOK, gin.H{
			"message": "ok",
			"data":    string(jsonTokens),
		})

}
