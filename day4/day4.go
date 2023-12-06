package day4

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

type Solution struct{}

type ScratchCard struct {
	Id      uint8
	Winners []uint8
	Numbers []uint8
}

func NewScratchCard(in []byte) *ScratchCard {
	line := bytes.NewBuffer(in)
	var sc ScratchCard

	n, err := fmt.Fscanf(line, "Card %d:", &sc.Id)
	if n == 0 {
		panic("Expected card")
	} else if err != nil {
		panic(err)
	}

	for {
		var v uint8
		n, err = fmt.Fscanf(line, "%d", &v)
		if n == 0 {
			break
		} else if err != nil {
			panic(err)
		}
		sc.Winners = append(sc.Winners, v)
	}

	_, err = fmt.Fscanf(line, "|")
	if err != nil {
		panic(err)
	}

	for {
		var v uint8
		n, err = fmt.Fscanf(line, "%d", &v)
		if n == 0 {
			break
		} else if err != nil {
			panic(err)
		}
		sc.Numbers = append(sc.Numbers, v)
	}

	return &sc
}

func (sc *ScratchCard) Value() int {
	points := 0
	for _, w := range sc.Winners {
		for _, n := range sc.Numbers {
			if w == n {
				if points == 0 {
					points = 1
				} else {
					points = points * 2
				}
			}
		}
	}
	return points
}

func (sc *ScratchCard) NWinners() int {
	count := 0
	for _, w := range sc.Winners {
		for _, n := range sc.Numbers {
			if w == n {
				count++
			}
		}
	}
	return count
}

func (s Solution) Part1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	sum := 0

	for scanner.Scan() {
		sc := NewScratchCard(scanner.Bytes())

		sum += sc.Value()
	}

	return sum
}

func (s Solution) Part2(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	cardsScored := 0
	var upcomingDuplicates []int

	for scanner.Scan() {
		sc := NewScratchCard(scanner.Bytes())
		var nCards = 1
		if len(upcomingDuplicates) > 0 {
			nCards += upcomingDuplicates[0]
			upcomingDuplicates = upcomingDuplicates[1:]
		}

		cardsWon := sc.NWinners()
		for i := 0; i < cardsWon; i++ {
			if len(upcomingDuplicates) < i+1 {
				upcomingDuplicates = append(upcomingDuplicates, 1*nCards)
			} else {
				upcomingDuplicates[i] += 1 * nCards
			}
		}

		cardsScored += nCards
	}

	return cardsScored
}
