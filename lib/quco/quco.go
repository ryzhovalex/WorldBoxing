package quco

import (
	"strings"
	"worldboxing/lib/quco/tokens"
	"worldboxing/lib/utils"
)

type Unit = map[string]any
type Response struct {
	Units []*Unit
	// How many objects were affected.
	// For create: number of created objects.
	// For get: length of `Units`.
	// For update: number of updated objects.
	// For delete: number of deleted objects.
	Affected int
}

// For now only GET is supported.
func Execute(query string) (*Response, *utils.Error) {
	return nil, nil
}

func lexical(query string) []*tokens.Token {
	result := []*tokens.Token{}
	buf := ""
	quoteOpened := false
	for _, x := range query {
		stringX := string(x)
		if quoteOpened && x == '"' {
			// Add string buffer as Name token, don't trim spaces.
			if len(buf) > 0 {
				result = append(result, &tokens.Token{
					Type:  tokens.Name,
					Value: buf,
				})
			}
			buf = ""
			// Add quote token.
			result = append(
				result,
				lexicalParseChunk(stringX),
			)
			quoteOpened = false
			continue
		}
		// If encounter instruction ending operators, we convert existing
		// buffer to token, empty the buffer and move on. But not for string
		// content.
		instructionEnding := x == '\n' || x == '(' || x == ')' || x == '"' || x == '=' || x == '.' || x == ',' || x == ' ' || x == '\t'
		if !quoteOpened && instructionEnding {
			// Write non-empty buffer.
			if len(strings.TrimSpace(buf)) > 0 {
				result = append(result, lexicalParseChunk(buf))
			}
			buf = ""
			result = append(
				result,
				lexicalParseChunk(stringX),
			)
			if x == '"' {
				quoteOpened = !quoteOpened
			}
			continue
		}
		buf += stringX
	}
	return result
}

func lexicalParseChunk(chunk string) *tokens.Token {
	var tokenType tokens.Type = tokens.Name

	tokenType, ok := tokens.TextToTokenType[chunk]
	if !ok {
		if utils.IsFloat(chunk) {
			tokenType = tokens.Float
		}
		if utils.IsInt(chunk) {
			tokenType = tokens.Integer
		}
	}

	return &tokens.Token{
		Type:  tokenType,
		Value: chunk,
	}
}

func interpretation(tokens []*tokens.Token) {

}
