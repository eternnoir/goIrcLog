package Database

import (
	Model "goIrcLog/Model"
	"strings"
    //"fmt"
)

type DbProvider interface {
	SaveMessage(msg *Model.Message) bool
	CheckTableExits() bool
	CreateTable() bool
}

func GetNewDbProvider(DbType, DbConnection string) DbProvider {
	DbType = strings.ToLower(DbType)
	var ret DbProvider
	switch DbType {
	case "sqlite":
		ret = NewSqliteDb(DbConnection)
	}
    ret.CreateTable()
	return ret
}
