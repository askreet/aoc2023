package day5

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"

	"github.com/askreet/aoc2023/advent"
	"github.com/askreet/aoc2023/day5/sparse_map"
)

type Solution struct{}

type Puzzle struct {
	Seeds []int
	Maps  []sparse_map.Map
}

func Parse(input io.Reader) *Puzzle {
	var p Puzzle
	scanner := bufio.NewScanner(input)
	scanner.Split(advent.ScanSections)

	haveReadHeader := false
	for scanner.Scan() {
		if !haveReadHeader {
			section := bytes.NewBuffer(scanner.Bytes()[6:])
			for {
				var seed int
				n, err := fmt.Fscanf(section, "%d", &seed)
				if n == 0 {
					break
				} else if err != nil {
					panic(err)
				}
				p.Seeds = append(p.Seeds, seed)
			}
			haveReadHeader = true
		} else {
			section := scanner.Bytes()
			idx := bytes.IndexByte(section, '\n')
			map_ := sparse_map.Parse(bytes.NewBuffer(section[idx+1:]))
			p.Maps = append(p.Maps, *map_)
		}
	}

	return &p
}

func (s Solution) Part1(input io.Reader) int {
	puzzle := Parse(input)

	lowestLocation := math.MaxInt
	for _, seed := range puzzle.Seeds {
		var value = seed
		for _, map_ := range puzzle.Maps {
			value = map_.Map(value)
		}
		if value < lowestLocation {
			lowestLocation = value
		}
	}

	return lowestLocation
}

func (s Solution) Part2(input io.Reader) int {
	puzzle := Parse(input)

	lowestLocation := math.MaxInt
	for i := 0; i < len(puzzle.Seeds); i += 2 {
		start := puzzle.Seeds[i]
		end := start + puzzle.Seeds[i+1]

		fmt.Println()
		for seed := start; seed < end; seed++ {
			if seed%100000 == 0 {
				fmt.Printf("\r[pair %d/%d] [seed %d/%d]", i/2, len(puzzle.Seeds)/2, (seed-start)+1, end-start)
			}
			var value = seed
			for _, map_ := range puzzle.Maps {
				value = map_.Map(value)
			}
			if value < lowestLocation {
				lowestLocation = value
			}
		}
	}

	return lowestLocation
}
