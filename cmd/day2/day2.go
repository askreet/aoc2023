package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	input, err := os.Open("in/day2.txt")
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}

	fmt.Println("part1: ", part1(input))

	_, _ = input.Seek(0, 0)
	fmt.Println("part2: ", part2(input))
}

func part1(in io.Reader) int {
	games, err := GamesFrom(in)
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

func part2(in io.Reader) int {
	games, err := GamesFrom(in)
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
