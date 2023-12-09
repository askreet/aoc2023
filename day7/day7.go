package day7

import (
	"bufio"
	"cmp"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
)

type Solution struct{}

func (s Solution) Calculate(input io.Reader, jokers bool) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var hands []Hand

	for scanner.Scan() {
		line := scanner.Text()
		d := strings.Split(line, " ")
		cards := d[0]
		bid, err := strconv.ParseInt(d[1], 10, 32)
		if err != nil {
			panic(err)
		}

		hands = append(hands, NewHand([]byte(cards), int(bid), jokers))
	}

	slices.SortFunc(hands, func(a, b Hand) int {
		return cmp.Compare(a.Value, b.Value)
	})

	total := 0
	for idx, hand := range hands {
		total += (idx + 1) * hand.Bid
		fmt.Printf("[nj=%d] rank=%d cards=%s value=%d type=%s\n", hand.nJokers(), idx+1, hand.Cards, hand.Value, hand.TypeName)
	}

	return total
}

func (s Solution) Part1(input io.Reader) int {
	return s.Calculate(input, false)
}

func (s Solution) Part2(input io.Reader) int {
	return s.Calculate(input, true)
}
