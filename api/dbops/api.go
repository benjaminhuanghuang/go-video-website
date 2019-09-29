package dbops

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func AddUserCredential(loginName string, pwd string) error {
	stmtInsert, err := dbConn.Prepare("INSERT INT users (login_name, pwd) VALUES (?, ?)")
	if err != nil {
		return err
	}
	stmtInsert.Exec(loginName, pwd)
	stmtInsert.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT users FROM users WHERE login_name=?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	var pwd string
	stmtOut.QueryRow(loginName).Scan(&pwd)
	stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE users FROM users WHERE login_name=? AND pwd=?")
	if err != nil {
		log.Printf("Delete user error: %s", err)
		return err
	}
	stmtDel.Exec(loginName, pwd)
	stmtDel.Close()
	return nil
}
