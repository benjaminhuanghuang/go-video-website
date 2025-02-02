package dbops

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// the global variable in the dops package
var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:rootpwd@tcp(localhost:3306)/video_server?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}
