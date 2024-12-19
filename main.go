package main

import (
	"os"
	"worldboxing/internal/cli"
	"worldboxing/lib/orwynn"
	"worldboxing/lib/utils"
)

func exit(ctx *cli.Context) *utils.Error {
	os.Exit(0)
	return nil
}

func main() {
	e := utils.LoadTranslationCsv("Static/Translations/en.csv", "en", ';')
	utils.Unwrap(e)
	cliTransport, e := cli.Init()
	utils.Unwrap(e)
	e = orwynn.Init(map[string]orwynn.Transport{
		"cli": cliTransport,
	})
	utils.Unwrap(e)

	cli.RegisterCommand("q", exit)
	cli.Start()
}
