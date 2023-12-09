package day9

import (
	"bufio"
	"io"
	"strconv"

	"github.com/johncgriffin/overflow"
)

type Solution struct{}

func (s Solution) Solve(input io.Reader, fn func(...int) int) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanBytes)

	var word []byte
	var numbers []int
	var total int
	for scanner.Scan() {
		b := scanner.Bytes()[0]
		switch b {
		case ' ':
			number, err := strconv.ParseInt(string(word), 10, 32)
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, int(number))
			word = word[0:0]

		case '\n':
			number, err := strconv.ParseInt(string(word), 10, 32)
			if err != nil {
				panic(err)
			}
			numbers = append(numbers, int(number))
			word = word[0:0]

			nn := fn(numbers...)
			nt, ok := overflow.Add(total, nn)
			if !ok {
				panic("overflow while calculating total")
			}
			total = nt
			numbers = numbers[0:0]

		default:
			word = append(word, b)
		}
	}
	if len(word) != 0 {
		panic("unprocessed result")
	}

	return total
}

func (s Solution) Part1(input io.Reader) int {
	return s.Solve(input, NextNumber)
}

func (s Solution) Part2(input io.Reader) int {
	return s.Solve(input, PrevNumber)
}

func NextNumber(numbers ...int) int {
	// Base case: all zeros.
	for _, n := range numbers {
		if n != 0 {
			goto NotZeros
		}
	}
	return 0

NotZeros:
	var diffs []int
	for i := 1; i < len(numbers); i++ {
		diffs = append(diffs, numbers[i]-numbers[i-1])
	}
	return numbers[len(numbers)-1] + NextNumber(diffs...)
}

func PrevNumber(numbers ...int) int {
	// Base case: all zeros.
	for _, n := range numbers {
		if n != 0 {
			goto NotZeros
		}
	}
	return 0

NotZeros:
	var diffs []int
	for i := 1; i < len(numbers); i++ {
		diffs = append(diffs, numbers[i]-numbers[i-1])
	}
	return numbers[0] - PrevNumber(diffs...)
}
