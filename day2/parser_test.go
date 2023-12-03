package day2

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestLex(t *testing.T) {
	// This program lacks a trailing newline on purpose to ensure end parsing.
	input := bytes.NewBufferString("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\n" +
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue")

	tokens, err := Lex(input)
	assert.NoError(t, err)

	expected := Tokens{
		{Ty: TGame},
		{Ty: TNumber, IntVal: 1},
		{Ty: TColon},
		{Ty: TNumber, IntVal: 3},
		{Ty: TColor, StrVal: "blue"},
		{Ty: TComma},
		{Ty: TNumber, IntVal: 4},
		{Ty: TColor, StrVal: "red"},
		{Ty: TSemicolon},
		{Ty: TNumber, IntVal: 1},
		{Ty: TColor, StrVal: "red"},
		{Ty: TComma},
		{Ty: TNumber, IntVal: 2},
		{Ty: TColor, StrVal: "green"},
		{Ty: TComma},
		{Ty: TNumber, IntVal: 6},
		{Ty: TColor, StrVal: "blue"},
		{Ty: TSemicolon},
		{Ty: TNumber, IntVal: 2},
		{Ty: TColor, StrVal: "green"},
		{Ty: TNewline},

		{Ty: TGame},
		{Ty: TNumber, IntVal: 2},
		{Ty: TColon},
		{Ty: TNumber, IntVal: 1},
		{Ty: TColor, StrVal: "blue"},
		{Ty: TComma},
		{Ty: TNumber, IntVal: 2},
		{Ty: TColor, StrVal: "green"},
		{Ty: TSemicolon},
		{Ty: TNumber, IntVal: 3},
		{Ty: TColor, StrVal: "green"},
		{Ty: TComma},
		{Ty: TNumber, IntVal: 4},
		{Ty: TColor, StrVal: "blue"},
		{Ty: TComma},
		{Ty: TNumber, IntVal: 1},
		{Ty: TColor, StrVal: "red"},
		{Ty: TSemicolon},
		{Ty: TNumber, IntVal: 1},
		{Ty: TColor, StrVal: "green"},
		{Ty: TComma},
		{Ty: TNumber, IntVal: 1},
		{Ty: TColor, StrVal: "blue"},
		{Ty: TEOF},
	}

	// Remove line context to aid in simpler matching.
	for i := range tokens {
		tokens[i].Line = 0
		tokens[i].Col = 0
	}

	assert.Equal(t, expected, tokens)
}

func TestLex_ErrorHandling(t *testing.T) {
	cases := []struct {
		Program  string
		ErrorMsg string
	}{
		{
			"Game 1 blue:",
			"unexpected TColon at line 1, col 12",
		},
	}

	for i, case_ := range cases {
		t.Run(fmt.Sprintf("program %d", i), func(t *testing.T) {
			input := bytes.NewBufferString(case_.Program)

			result, err := Lex(input)
			spew.Dump(result)
			assert.ErrorContains(t, err, case_.ErrorMsg)
		})
	}
}

func TestParse(t *testing.T) {
	tokens := Tokens{
		{Ty: TGame},
		{Ty: TNumber, IntVal: 1},
		{Ty: TColon},
		{Ty: TNumber, IntVal: 3},
		{Ty: TColor, StrVal: "blue"},
		{Ty: TComma},
		{Ty: TNumber, IntVal: 4},
		{Ty: TColor, StrVal: "red"},
		{Ty: TSemicolon},
		{Ty: TNumber, IntVal: 1},
		{Ty: TColor, StrVal: "red"},
		{Ty: TComma},
		{Ty: TNumber, IntVal: 2},
		{Ty: TColor, StrVal: "green"},
		{Ty: TComma},
		{Ty: TNumber, IntVal: 6},
		{Ty: TColor, StrVal: "blue"},
		{Ty: TSemicolon},
		{Ty: TNumber, IntVal: 2},
		{Ty: TColor, StrVal: "green"},
		{Ty: TEOF},
	}

	games, err := Parse(tokens)
	assert.NoError(t, err)

	expected := []Game{
		{
			Id: 1,
			Sets: [][]Reveal{
				{
					Reveal{Color: "blue", Number: 3},
					Reveal{Color: "red", Number: 4},
				},
				{
					Reveal{Color: "red", Number: 1},
					Reveal{Color: "green", Number: 2},
					Reveal{Color: "blue", Number: 6},
				},
				{
					Reveal{Color: "green", Number: 2},
				},
			},
		},
	}
	assert.Equal(t, expected, games)
}
