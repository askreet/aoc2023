package day1

import (
	"bufio"
	"io"

	"go.arsenm.dev/pcre"
)

type Solution struct{}

func (s Solution) Part1(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		total += (ldigit(line) * 10) + rdigit(line)
	}

	return total
}

func (s Solution) Part2(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()

		digits := convertCaptureGroupsToDigits(PartTwoRegex.FindAllStringSubmatchIndex(line, -1))

		ldigit := digits[0]
		rdigit := digits[len(digits)-1]

		total += (ldigit * 10) + rdigit
	}

	return total
}

func ldigit(in string) int {
	runes := []rune(in)

	for i := 0; i < len(runes); i++ {
		if runes[i] >= 48 && runes[i] <= 57 {
			return int(runes[i]) & 0b1111
		}
	}

	panic("expected digit in input value: " + in)
}

func rdigit(in string) int {
	runes := []rune(in)

	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] >= 48 && runes[i] <= 57 {
			return int(runes[i]) & 0b1111
		}
	}

	panic("expected digit in input value: " + in)
}

// https://xkcd.com/208/
var PartTwoRegex = pcre.MustCompile(
	"(0)" +
		"|(1|o(?=ne))" +
		"|(2|t(?=wo))" +
		"|(3|t(?=hree))" +
		"|(4|f(?=our))" +
		"|(5|f(?=ive))" +
		"|(6|s(?=ix))" +
		"|(7|s(?=even))" +
		"|(8|e(?=ight))" +
		"|(9|n(?=ine))")

func convertCaptureGroupsToDigits(indexSets [][]int) []int {
	var digits []int

	for _, indexSet := range indexSets {
		// indexSets are pairs of capture groups that were found. The first
		// set is the full match, which we ignore. From there, we can infer the digit
		// from which capture group is not missing (-1:-1).
		for start := 2; start < len(indexSet); start += 2 {
			if indexSet[start] != -1 {
				foundDigit := (start / 2) - 1
				digits = append(digits, foundDigit)
				goto NextMatch
			}
		}

	NextMatch:
	}

	return digits
}
