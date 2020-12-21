package controller

import (
	"database/sql"

	"github.com/aybjax/go-gin-test/data_structures"
	"github.com/aybjax/go-gin-test/database"
)

// GetData -> get data from database
func GetData(data *data_structures.GetData) []data_structures.TokenData {
	var stmt *sql.Stmt
	var rows *sql.Rows
	var err error

	if len(data.Device) > 0 {
		stmt = database.GetDeviceStmt()
		rows, err = stmt.Query(data.Device)
	} else {
		stmt = database.GetUserStmt()
		rows, err = stmt.Query(data.User)
	}

	if err != nil {
		// fmt.Println(err.Error())
		return nil
	}

	tokens := make([]data_structures.TokenData, 0)
	var num int
	for rows.Next() {
		var token data_structures.TokenData
		rows.Scan(
			&num,
			&token.User,
			&token.Device,
			&token.Token,
			&token.Os,
			&token.Version,
		)
		tokens = append(tokens, token)
	}

	return tokens
}
