package main

import (
	"os"
	"worldboxing/internal/cli"
	"worldboxing/lib/database"
	"worldboxing/lib/quco"
	"worldboxing/lib/utils"
)

func qucoGet(ctx *cli.Context) *utils.Error {
	_, e := quco.Execute(ctx.Call.Command)
	return e
}

func write(_ *cli.Context) *utils.Error {
	be := database.Tx.Commit()
	if be != nil {
		return utils.DefaultError()
	}
	return nil
}

func rollback(_ *cli.Context) *utils.Error {
	be := database.Tx.Rollback()
	if be != nil {
		return utils.DefaultError()
	}
	// restart transaction after rollback
	database.BeginGlobalTransaction()
	return nil
}

func quit(ctx *cli.Context) *utils.Error {
	os.Exit(0)
	return nil
}

func main() {
	var e *utils.Error

	e = utils.LoadTranslationCsv("Translations/en.csv", "en", ';')
	utils.Unwrap(e)
	e = database.Init()
	utils.Unwrap(e)
	e = database.BeginGlobalTransaction()
	utils.Unwrap(e)

	cli.RegisterCommand("get", qucoGet)
	cli.RegisterCommand("write", write)
	cli.RegisterCommand("rollback", rollback)
	cli.RegisterCommand("quit", quit)
	cli.Start()
}
