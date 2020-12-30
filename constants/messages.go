package constants

import "github.com/gin-gonic/gin"

// INCORRECT_DATA -> gin.H type for response
var INCORRECT_DATA = gin.H{
	"message": "incorrect data",
}

// DATABASE_ERROR -> gin.H type for response
var DATABASE_ERROR = gin.H{
	"message": "database error",
}

// STATUS_OK -> status code of ok
var STATUS_OK = gin.H{
	"message": "ok",
}
