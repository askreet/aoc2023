package day2

import "io"

type Game struct {
	Id   int
	Sets [][]Reveal
}

type Games []Game

func GamesFrom(in io.Reader) (Games, error) {
	tokens, err := Lex(in)
	if err != nil {
		return nil, err
	}

	return Parse(tokens)
}
