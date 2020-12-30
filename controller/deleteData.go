package controller

import (
	"github.com/aybjax/go-gin-test/data_structures"
	"github.com/aybjax/go-gin-test/database"
)

// DeleteData -> delete data from database
func DeleteData(data *data_structures.DeleteData) bool {
	stmt := database.DeleteStmt()

	_, err := stmt.Exec(data.Token)

	if err != nil {
		return false
	}

	return true
}
