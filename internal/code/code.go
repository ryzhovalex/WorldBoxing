package code

import "worldboxing/lib/utils"

type Code = utils.Code

const (
	Ok                               Code = 0
	Error                            Code = 1
	CliCallParsingError                   = 2
	CliAlreadyRegisteredCommandError      = 3
	CliMissingCommandError                = 4
)
