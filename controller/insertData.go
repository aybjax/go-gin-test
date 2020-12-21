package controller

import (
	"github.com/aybjax/go-gin-test/data_structures"
	"github.com/aybjax/go-gin-test/database"
)

// InsertData -> insert data into db
func InsertData(data data_structures.InsertData) bool {
	// user_id, device_id, token, os, version
	stmt := database.InsertStmt()

	_, err := stmt.Exec(
		data.User,
		data.Device,
		data.Token,
		data.Os,
		data.Version,
	)

	if err != nil {
		// fmt.Println(err.Error())
		return false
	}

	return true

}
