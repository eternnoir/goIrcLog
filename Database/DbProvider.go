package Database

import(
   Model "goIrcLog/Model"
)

type DbProvider interface {
    SaveMessage(msg *Model.Message) bool
    CheckTableExits() bool
    CreateTable() bool
}
