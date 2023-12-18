package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/askreet/aoc2023/advent"
	"github.com/askreet/aoc2023/day1"
	"github.com/askreet/aoc2023/day10"
	"github.com/askreet/aoc2023/day11"
	"github.com/askreet/aoc2023/day12"
	"github.com/askreet/aoc2023/day13"
	"github.com/askreet/aoc2023/day14"
	"github.com/askreet/aoc2023/day2"
	"github.com/askreet/aoc2023/day3"
	"github.com/askreet/aoc2023/day4"
	"github.com/askreet/aoc2023/day5"
	"github.com/askreet/aoc2023/day6"
	"github.com/askreet/aoc2023/day7"
	"github.com/askreet/aoc2023/day8"
	"github.com/askreet/aoc2023/day9"
)

var Days = []advent.Interface{
	day1.Solution{},
	day2.Solution{},
	day3.Solution{},
	day4.Solution{},
	day5.Solution{},
	day6.Solution{},
	day7.Solution{},
	day8.Solution{},
	day9.Solution{},
	day10.Solution{},
	day11.Solution{},
	day12.Solution{},
	day13.Solution{},
	day14.Solution{},
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

	var part = ""
	if len(os.Args) >= 3 {
		part = os.Args[2]
	}

	if part == "" || part == "1" {
		fmt.Println("part1: ", Days[day-1].Part1(input))
	}

	_, _ = input.Seek(0, 0)
	if part == "" || part == "2" {
		fmt.Println("part2: ", Days[day-1].Part2(input))
	}
}
