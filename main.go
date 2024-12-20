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

func exit(ctx *cli.Context) *utils.Error {
	os.Exit(0)
	return nil
}

func main() {
	var e *utils.Error

	e = utils.LoadTranslationCsv("Static/Translations/en.csv", "en", ';')
	utils.Unwrap(e)
	e = database.Init()
	utils.Unwrap(e)

	cli.RegisterCommand("g", qucoGet)
	cli.RegisterCommand("q", exit)
	cli.Start()
}
