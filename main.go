package main

import (
	"os"
	"worldboxing/internal/cli"
	"worldboxing/lib/database"
	"worldboxing/lib/utils"
)

func qucoGet(ctx *cli.Context) *utils.Error {
	return nil
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
