package code

import "worldboxing/lib/utils"

type Code = utils.Code

const Ok Code = utils.CodeOk
const Error Code = utils.CodeError

// CLI
const CliCallParsing Code = 2
const CliCommandAlreadyRegistered Code = 3
const CliNoSuchCommand Code = 4
