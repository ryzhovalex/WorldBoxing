package quco

import "testing"

func TestLexicalOk(t *testing.T) {
	tokens := lexical(`GET
firstname="John"
surname="Marlow"
country.IN=("Demacia", "Imladris")
`)
	for i, token := range tokens {
		print(i, ": ", token.Type, " ", token.Value, "\n")
	}
	t.Fail()
}
