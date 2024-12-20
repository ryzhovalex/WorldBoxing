package main

import (
	"os"
	"worldboxing/internal/cli"
	"worldboxing/internal/sim"
	"worldboxing/lib/database"
	"worldboxing/lib/quco"
	"worldboxing/lib/utils"
)

func qucoGet(ctx *cli.Context) *utils.Error {
	_, e := quco.Execute(ctx.Call.Command)
	return e
}

func simStart(ctx *cli.Context) *utils.Error {
	return sim.Start()
}

func write(_ *cli.Context) *utils.Error {
	be := database.T.Commit()
	if be != nil {
		return utils.DefaultError("")
	}
	return nil
}

func rollback(_ *cli.Context) *utils.Error {
	be := database.T.Rollback()
	if be != nil {
		return utils.DefaultError("")
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

	cli.RegisterCommand("g", qucoGet)
	cli.RegisterCommand("w", write)
	cli.RegisterCommand("r", rollback)
	cli.RegisterCommand("q", quit)

	cli.RegisterCommand("sim.start", simStart)

	cli.Start()
}
