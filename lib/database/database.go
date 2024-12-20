package database

import (
	"worldboxing/lib/utils"

	_ "github.com/glebarez/go-sqlite"
	"github.com/jmoiron/sqlx"
)

var D *sqlx.DB
var T *sqlx.Tx

func Init() *utils.Error {
	var baseError error
	D, baseError = sqlx.Connect("sqlite", "./Var/Main.db")
	if baseError != nil {
		return utils.DefaultError("")
	}
	return nil
}

func BeginGlobalTransaction() *utils.Error {
	var be error
	T, be = D.Beginx()
	if be != nil {
		return utils.DefaultError("")
	}
	return nil
}
