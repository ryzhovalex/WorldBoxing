package database

import (
	"worldboxing/lib/utils"

	_ "github.com/glebarez/go-sqlite"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func Init() *utils.Error {
	var baseError error
	Db, baseError = sqlx.Connect("sqlite", "./Var/Main.db")
	if baseError != nil {
		return utils.DefaultError()
	}
	return nil
}
