package quco

import (
	"fmt"
	"strings"
	"worldboxing/lib/database"
	"worldboxing/lib/quco/tokens"
	"worldboxing/lib/utils"
)

type Code = utils.Code

const (
	CodeUnsupportedAction      Code = 1
	CodeCollectionParsing      Code = 2
	CodeEmptyActionBody        Code = 3
	CodeSqlQueryExecutionError Code = 4
	CodeTokenParsingError      Code = 5
)

type Unit = map[string]any
type Response struct {
	Units *[]Unit
	// How many objects were affected.
	// For create: number of created objects.
	// For get: length of `Units`.
	// For update: number of updated objects.
	// For delete: number of deleted objects.
	Affected int
}

// For now only GET is supported.
func Execute(query string) (*Response, *utils.Error) {
	t := lexical(query)

	dbQuery, e := interpretation(t)
	if e != nil {
		return nil, e
	}

	units := &[]Unit{}
	be := database.Tx.Select(units, dbQuery, nil)
	if be != nil {
		return nil, utils.NewError(CodeSqlQueryExecutionError, "")
	}

	return &Response{
		Units: units,
		// TODO: change for non-get operations
		Affected: len(*units),
	}, nil
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
		floatBuilding := x == '.' && utils.IsInt(buf)
		if !floatBuilding && !quoteOpened && instructionEnding {
			// Write non-empty buffer.
			if len(utils.RemoveSpaces(buf)) > 0 {
				result = append(result, lexicalParseChunk(buf))
			}
			buf = ""
			// Whitespaces are ignored everywhere except inside
			// strings.
			if x != ' ' && x != '\t' {
				result = append(
					result,
					lexicalParseChunk(stringX),
				)
			}
			if x == '"' {
				quoteOpened = !quoteOpened
			}
			continue
		}
		buf += stringX
	}
	// Save rest of the buffer as token if it's not empty.
	if len(utils.RemoveSpaces(buf)) > 0 {
		result = append(result, lexicalParseChunk(buf))
	}
	if result[len(result)-1].Type != tokens.EndInstruction {
		result = append(result, &tokens.Token{Type: tokens.EndInstruction, Value: "\n"})
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

type QucoAction int

const (
	ActionGet QucoAction = iota
	ActionCreate
	ActionSet
	ActionUpdate
	ActionDelete
)

func interpretation(lexicalTokens []*tokens.Token) (string, *utils.Error) {
	if len(lexicalTokens) == 0 {
		return "", nil
	}

	var action QucoAction
	switch lexicalTokens[0].Type {
	case tokens.Get:
		action = ActionGet
	default:
		return "", utils.NewError(CodeUnsupportedAction, "")
	}

	if len(lexicalTokens) < 2 || lexicalTokens[1].Type != tokens.Name {
		return "", utils.NewError(CodeCollectionParsing, "")
	}
	collection := lexicalTokens[1].Value

	if len(lexicalTokens) < 3 {
		return "", utils.NewError(CodeEmptyActionBody, "")
	}

	switch action {
	case ActionGet:
		// Body tokens start after Action, Collection and EndInstruction
		// tokens.
		return get(collection, lexicalTokens[3:])
	default:
		return "", utils.NewError(CodeUnsupportedAction, "")
	}
}

func get(
	collection string,
	bodyTokens []*tokens.Token,
) (string, *utils.Error) {
	whereQuery, e := generateWhereQuery(bodyTokens)
	if e != nil {
		return "", e
	}
	dbQuery := fmt.Sprintf(`
SELECT * FROM %s
WHERE %s
`, collection, whereQuery)
	return dbQuery, nil
}

func generateWhereQuery(bodyTokens []*tokens.Token) (string, *utils.Error) {
	conditions := []string{}
	instructionKey := ""
	instructionValue := ""
	// By default all instructions are assignment (or equality in case of GET).
	instructionComparisonType := tokens.Assignment
	instructionComparisonTypeSet := false
	dotTokenLast := false
	assignmentTokenAppeared := false

	for _, t := range bodyTokens {
		// Compose left-hand instruction key from dot pieces.
		if t.Type == tokens.Name && !assignmentTokenAppeared {
			if instructionKey != "" {
				instructionKey += "." + t.Value
			} else {
				instructionKey = t.Value
			}
		}
		if tokens.IsComparisonToken(t) {
			if instructionComparisonTypeSet && t.Type != tokens.Assignment {
				return "", utils.NewError(CodeTokenParsingError, "Duplicate comparison instruction.")
			}
			if instructionComparisonTypeSet && instructionComparisonType == t.Type && t.Type == tokens.Assignment {
				return "", utils.NewError(CodeTokenParsingError, "Duplicate assignment instruction.")
			}
			if len(instructionKey) == 0 {
				return "", utils.NewError(CodeTokenParsingError, "Cannot assign comparsion token without active instruction.")
			}
			// Assignment token can go without preceding dot.
			if !dotTokenLast && t.Type != tokens.Assignment {
				return "", utils.NewError(CodeTokenParsingError, "Comparison token without dot operator.")
			}
			if assignmentTokenAppeared {
				return "", utils.NewError(CodeTokenParsingError, "Invalid assignment token placement.")
			}
			// Avoid resetting of true comparison type by the required
			// assignment token.
			if !instructionComparisonTypeSet {
				instructionComparisonTypeSet = true
				instructionComparisonType = t.Type
			}
		} else {
			instructionComparisonTypeSet = false
		}
		if t.Type == tokens.Assignment {
			if assignmentTokenAppeared {
				return "", utils.NewError(CodeTokenParsingError, "Duplicate assignment token.")
			}
			if instructionKey == "" {
				return "", utils.NewError(CodeTokenParsingError, "Invalid instruction target.")
			}
			assignmentTokenAppeared = true
			continue
		}
		if t.Type == tokens.EndInstruction {
			if instructionKey == "" || instructionValue == "" || !assignmentTokenAppeared {
				return "", utils.NewError(CodeTokenParsingError, "Invalid instruction.")
			}

			condition, e := parseInstructionToCondition(instructionKey, instructionComparisonType, instructionValue)
			if e != nil {
				return "", e
			}
			conditions = append(conditions, condition)

			instructionKey = ""
			instructionValue = ""
			assignmentTokenAppeared = false
			instructionComparisonTypeSet = false
			dotTokenLast = false
			continue
		}
		if assignmentTokenAppeared {
			// Treat the following tokens as the right side of the instruction.
			if (t.Type == tokens.ContainerOpen || t.Type == tokens.ContainerClose || t.Type == tokens.Comma) && instructionComparisonType != tokens.In {
				return "", utils.NewError(CodeTokenParsingError, "Containers are allowed only for IN comparison.")
			}
			prevInstructionValue := instructionValue
			// These types converted to SQL "as is" in Quco.
			if t.Type == tokens.ContainerOpen || t.Type == tokens.ContainerClose || t.Type == tokens.Comma || t.Type == tokens.Integer || t.Type == tokens.Float || t.Type == tokens.Name {
				instructionValue += t.Value
			}
			if t.Type == tokens.Quote {
				instructionValue += "'"
			}
			if t.Type == tokens.True {
				instructionValue += "1"
			}
			if t.Type == tokens.False {
				instructionValue += "0"
			}

			// No changes => error.
			if prevInstructionValue == instructionValue {
				return "", utils.NewError(CodeTokenParsingError, "Right-side token produced no adjustments to instruction value.")
			}
		}

		dotTokenLast = t.Type == tokens.Dot
	}

	return strings.Join(conditions, " AND "), nil
}

func parseInstructionToCondition(
	instructionKey string,
	comparisonType tokens.Type,
	instructionValue string,
) (string, *utils.Error) {
	comparisonText := ""
	prefix := ""
	switch comparisonType {
	case tokens.Assignment:
		comparisonText = "="
	case tokens.Ne:
		comparisonText = "="
		prefix = "NOT "
	case tokens.In:
		comparisonText = "IN"
	case tokens.Lt:
		comparisonText = "<"
	case tokens.Le:
		comparisonText = "<="
	case tokens.Gt:
		comparisonText = ">"
	case tokens.Ge:
		comparisonText = ">="
	default:
		return "", utils.NewError(CodeTokenParsingError, "Invalid comparison type.")
	}
	condition := fmt.Sprintf("%s%s %s %s", prefix, instructionKey, comparisonText, instructionValue)
	return condition, nil
}
