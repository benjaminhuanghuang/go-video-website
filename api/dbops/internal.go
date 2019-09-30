package dbops

import (
	"database/sql"
	"log"
	"strconv"
	"sync"

	"../defs"
)

func InsertSession(sid string, ttl int64, login_name string) error {
	ttlstr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("INSERT INTO session (session_id, TTL, login) VALUES (?,?,?)")

	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(sid, ttlstr, login_name)
	if err != nil {
		return nil
	}
	defer stmtIns.Close()
	return nil
}

func RetrieveSesson(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare("SELECT TTL,login_name from sessions WHER session_id=?")
	if err != nil {
		return nil, err
	}

	var ttl string
	var uname string
	err = stmtOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		ss.TTL = res
		ss.Username = uname
	} else {
		return nil, err
	}
	defer stmtOut.Close()
	return ss, nil
}

func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmtOut, err := dbConn.Prepare("SELECT * FROM sessions")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	rows, err := stmtOut.Query()
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	for rows.Next() {
		var id string
		var ttlstr string
		var login_name string
		if err := rows.Scan(&id, &ttlstr, &login_name); err != nil {
			log.Printf("retrive sessions error: %s", err)
			break
		}
		if ttl, errl := strconv.ParseInt(ttlstr, 10, 64); errl == nil {
			ss := &defs.SimpleSession{Username: login_name, TTL: ttl}
			m.Store(id, ss)
			log.Printf("Session id: %s, ttl: %d", id, ss.TTL)
		}
	}
	return m, nil
}

func DeleteSession(sid string) error {
	stmtOut, err := dbConn.Prepare("DELETE FROM sessions WHERE session_id=?")
	if err != nil {
		log.Printf("%s", err)
		return err
	}

	if _, err := stmtOut.Exec(sid); err != nil {
		return err
	}
	return nil
}
