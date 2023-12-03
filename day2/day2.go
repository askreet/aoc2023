package day2

import (
	"io"
)

type Solution struct{}

func (s Solution) Part1(input io.Reader) int {
	games, err := GamesFrom(input)
	if err != nil {
		panic(err)
	}

	bag := NewBag()
	bag.AddN("red", 12)
	bag.AddN("green", 13)
	bag.AddN("blue", 14)

	possibleGameIDSum := 0
	for _, game := range games {
		for _, set := range game.Sets {
			for _, reveal := range set {
				if !bag.CouldReveal(&reveal) {
					goto NextGame
				}
			}
		}
		possibleGameIDSum += game.Id

	NextGame:
	}

	return possibleGameIDSum
}

func (s Solution) Part2(input io.Reader) int {
	games, err := GamesFrom(input)
	if err != nil {
		panic(err)
	}

	total := 0
	for _, game := range games {
		bag := MinBagFor(&game)
		total += bag.Power()
	}
	return total
}
