package main

import (
	"os"
	"worldboxing/internal/cli"
	"worldboxing/lib/utils"
)

func exit(ctx *cli.Context) *utils.Error {
	os.Exit(0)
	return nil
}

func main() {
	e := utils.LoadTranslationCsv("Static/Translations/en.csv", "en", ';')
	utils.Unwrap(e)

	cli.RegisterCommand("q", exit)
	cli.Start()
}
