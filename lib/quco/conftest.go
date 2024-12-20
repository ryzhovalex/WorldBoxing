package quco

import (
	"testing"
	"worldboxing/lib/quco/tokens"
)

func compareTokens(t *testing.T, expectedTokens []*tokens.Token, realTokens []*tokens.Token) {
	if len(realTokens) != len(expectedTokens) {
		t.FailNow()
	}
	for i, token := range realTokens {
		if expectedTokens[i].Type != token.Type {
			t.FailNow()
		}
		if expectedTokens[i].Value != token.Value {
			t.FailNow()
		}
	}
}
