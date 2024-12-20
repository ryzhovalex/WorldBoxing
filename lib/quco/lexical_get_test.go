package quco

import (
	"testing"
	"worldboxing/lib/quco/tokens"
)

func TestStringOk(t *testing.T) {
	realTokens := lexical(`GET Person
Name="GET"`)
	compareTokens(
		t,
		[]*tokens.Token{
			{
				Type:  tokens.Get,
				Value: "GET",
			},
			{
				Type:  tokens.Name,
				Value: "Person",
			},
			{
				Type:  tokens.EndInstruction,
				Value: "\n",
			},
			{
				Type:  tokens.Name,
				Value: "Name",
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
			{
				Type:  tokens.EndInstruction,
				Value: "\n",
			},
		},
		realTokens,
	)
}

func TestIntegerOk(t *testing.T) {
	realTokens := lexical(`GET Person
Age=100`)
	compareTokens(
		t,
		[]*tokens.Token{
			{
				Type:  tokens.Get,
				Value: "GET",
			},
			{
				Type:  tokens.Name,
				Value: "Person",
			},
			{
				Type:  tokens.EndInstruction,
				Value: "\n",
			},
			{
				Type:  tokens.Name,
				Value: "Age",
			},
			{
				Type:  tokens.Assignment,
				Value: "=",
			},
			{
				Type:  tokens.Integer,
				Value: "100",
			},
			{
				Type:  tokens.EndInstruction,
				Value: "\n",
			},
		},
		realTokens,
	)
}

func TestContainerOk(t *testing.T) {
	realTokens := lexical(`GET Person
Salary.IN=(100, 200)`)
	compareTokens(
		t,
		[]*tokens.Token{
			{
				Type:  tokens.Get,
				Value: "GET",
			},
			{
				Type:  tokens.Name,
				Value: "Person",
			},
			{
				Type:  tokens.EndInstruction,
				Value: "\n",
			},
			{
				Type:  tokens.Name,
				Value: "Salary",
			},
			{
				Type:  tokens.Dot,
				Value: ".",
			},
			{
				Type:  tokens.In,
				Value: "IN",
			},
			{
				Type:  tokens.Assignment,
				Value: "=",
			},
			{
				Type:  tokens.ContainerOpen,
				Value: "(",
			},
			{
				Type:  tokens.Integer,
				Value: "100",
			},
			{
				Type:  tokens.Comma,
				Value: ",",
			},
			{
				Type:  tokens.Integer,
				Value: "200",
			},
			{
				Type:  tokens.ContainerClose,
				Value: ")",
			},
			{
				Type:  tokens.EndInstruction,
				Value: "\n",
			},
		},
		realTokens,
	)
}

func TestFloatOk(t *testing.T) {
	realTokens := lexical(`GET Person
Salary=10.5`)
	compareTokens(
		t,
		[]*tokens.Token{
			{
				Type:  tokens.Get,
				Value: "GET",
			},
			{
				Type:  tokens.Name,
				Value: "Person",
			},
			{
				Type:  tokens.EndInstruction,
				Value: "\n",
			},
			{
				Type:  tokens.Name,
				Value: "Salary",
			},
			{
				Type:  tokens.Assignment,
				Value: "=",
			},
			{
				Type:  tokens.Float,
				Value: "10.5",
			},
			{
				Type:  tokens.EndInstruction,
				Value: "\n",
			},
		},
		realTokens,
	)
}
