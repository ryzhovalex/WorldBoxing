package quco

import (
	"testing"
	"worldboxing/lib/quco/tokens"
)

func TestStringOk(t *testing.T) {
	realTokens := lexical(`GET
name="GET"`)
	compareTokens(
		t,
		[]*tokens.Token{
			{
				Type:  tokens.Get,
				Value: "GET",
			},
			{
				Type:  tokens.Name,
				Value: "name",
			},
			{
				Type:  tokens.Assignment,
				Value: "=",
			},
			{
				Type:  tokens.Quote,
				Value: "\"",
			},
			{
				Type:  tokens.Name,
				Value: "GET",
			},
			{
				Type:  tokens.Quote,
				Value: "\"",
			},
		},
		realTokens,
	)
}

func TestIntegerOk(t *testing.T) {
	realTokens := lexical(`GET
age=100`)
	compareTokens(
		t,
		[]*tokens.Token{
			{
				Type:  tokens.Get,
				Value: "GET",
			},
			{
				Type:  tokens.Name,
				Value: "age",
			},
			{
				Type:  tokens.Assignment,
				Value: "=",
			},
			{
				Type:  tokens.Integer,
				Value: "100",
			},
		},
		realTokens,
	)
}
