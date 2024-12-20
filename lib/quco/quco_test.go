package quco

import "testing"

func TestLexicalOk(t *testing.T) {
	tokens := lexical(`GET
name="GET"
comment="Average height. Average weight. IN(the mood)."
country.IN=("Demacia", "Imladris")
`)
	print("---\n")
	for i, token := range tokens {
		print(i, ": ", token.Type, " ", "`", token.Value, "`", "\n")
	}
	t.Fail()
}
