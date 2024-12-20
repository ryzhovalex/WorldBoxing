package code

import "worldboxing/lib/utils"

type Code = utils.Code

const (
	Hello      = -1
	Ok    Code = 0
	Error      = 1

	// CLI
	CliCallParsing              = 2
	CliCommandAlreadyRegistered = 3
	CliNoSuchCommand            = 4
)
