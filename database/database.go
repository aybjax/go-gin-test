package database

import (
	"database/sql"
	"fmt"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var createTableNotExist *sql.Stmt
var insertData *sql.Stmt
var getDataByUser *sql.Stmt
var getDataByDevice *sql.Stmt
var deleteData *sql.Stmt

// InsertStmt -> get statement for insertion
func InsertStmt() *sql.Stmt {
	return insertData
}

// DeleteStmt -> get statement for insertion
func DeleteStmt() *sql.Stmt {
	return deleteData
}

// GetUserStmt -> get statement for get
func GetUserStmt() *sql.Stmt {
	return getDataByUser
}

// GetDeviceStmt -> get statement for get
func GetDeviceStmt() *sql.Stmt {
	return getDataByDevice
}

// Close should be called as defer in main()
func Close() {
	db.Close()
	fmt.Println("Database closed")
}

func errorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// establish connection and init Stmts
func init() {
	if db != nil {
		return
	}
	var err error
	db, err = sql.Open("mysql", fmt.Sprintf(
		"%s:%s@/%s",
		"root",
		"",
		"test_go",
	))
	errorCheck(err)

	createTables()

	prepareStatements()
}

// create table if not exist
func createTables() {
	var err error
	createTableNotExist, err = db.Prepare(`
		CREATE TABLE IF NOT EXISTS tokens(
			id int primary key auto_increment,
			user_id int,
			device_id VARCHAR(50),
			token VARCHAR(50),
			os VARCHAR(10),
			version VARCHAR(10))
	`)
	errorCheck(err)
	createTableNotExist.Exec()
}

// prepare Stmts
func prepareStatements() {
	var err error

	insertData, err = db.Prepare(`
		INSERT INTO tokens(user_id, device_id, token, version, os)
		VALUES (?, ?, ?, ?, ?)
	`)
	errorCheck(err)

	getDataByUser, err = db.Prepare("SELECT * FROM tokens WHERE user_id=?;")
	errorCheck(err)

	getDataByDevice, err = db.Prepare("SELECT * FROM tokens WHERE device_id=?;")
	errorCheck(err)

	deleteData, err = db.Prepare("DELETE FROM tokens WHERE token=?;")
	errorCheck(err)
}
