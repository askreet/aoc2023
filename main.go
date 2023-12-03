package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/askreet/aoc2023/advent"
	"github.com/askreet/aoc2023/day1"
	"github.com/askreet/aoc2023/day2"
	"github.com/askreet/aoc2023/day3"
)

var Days = []advent.Interface{
	day1.Solution{},
	day2.Solution{},
	day3.Solution{},
}

func main() {
	day, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}

	input, err := os.Open(fmt.Sprintf("in/day%d.txt", day))
	if err != nil {
		fmt.Println("err: ", err)
		os.Exit(1)
	}

	fmt.Println("part1: ", Days[day-1].Part1(input))

	_, _ = input.Seek(0, 0)
	fmt.Println("part2: ", Days[day-1].Part2(input))
}
