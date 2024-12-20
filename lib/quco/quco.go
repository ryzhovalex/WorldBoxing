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
	for _, x := range query {
		// If encounter instruction ending operators, we convert existing
		// buffer to token, empty the buffer and move on.
		if x == '\n' || x == '(' || x == ')' || x == '"' {
			stringX := string(x)
			result = append(
				result,
				lexicalParseChunk(stringX),
			)
			if len(strings.TrimSpace(buf)) > 0 {
				result = append(result, lexicalParseChunk(buf))
			}
			buf = ""
			continue
		}
		buf += string(x)
	}
	return result
}

func lexicalParseChunk(chunk string) *tokens.Token {
	var tokenType tokens.Type = tokens.Name

	chunk = strings.TrimSpace(chunk)
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
