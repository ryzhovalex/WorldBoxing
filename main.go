package main

import (
	"worldboxing/internal/cli"
	"worldboxing/lib/utils"
)

func main() {
	e := utils.LoadTranslationCsv("Static/Translation.en.csv", "en", ';')
	utils.Unwrap(e)
	cli.Start()
}
