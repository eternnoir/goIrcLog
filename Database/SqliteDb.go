package Database

import (
	sqlite3 "code.google.com/p/go-sqlite/go1/sqlite3"
	"fmt"
	Model "goIrcLog/Model"
)

type SqlLiteDb struct {
	DbPath string
}

func (db *SqlLiteDb) SaveMessage(msg *Model.Message) bool {
	return false
}

func (db *SqlLiteDb) CheckTableExits() bool {
	return false
}
func (db *SqlLiteDb) CreateTable() bool {
	sqlStr :=
	`
    create table if not exists LogMsg(Message TEXT,Nick TEXT,Channel TEXT,Time INTEGER)
    `
	conn := db.getConn()
	err := conn.Exec(sqlStr)
	if err != nil {
		fmt.Println("Can not create table")
		return false
	} else {
		return true
	}
}

func (db *SqlLiteDb) getConn() *sqlite3.Conn {
	c, _ := sqlite3.Open(db.DbPath)
	return c
}
