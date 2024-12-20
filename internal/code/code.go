package code

import "worldboxing/lib/utils"

type Code = utils.Code

const (
	Ok                               Code = 0
	Error                            Code = 1
	CliCallParsingError              Code = 2
	CliAlreadyRegisteredCommandError Code = 3
	CliMissingCommandError           Code = 4
)
