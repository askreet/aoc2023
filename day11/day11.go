package day11

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"io"

	"github.com/askreet/aoc2023/lib"
)

type Solution struct{}

func (s Solution) Solve(input io.Reader, amount int) int {
	buf, _ := io.ReadAll(input)
	fmt.Println("found", bytes.Count(buf, []byte{'#'}), "galaxies")

	scanner := bufio.NewScanner(bytes.NewBuffer(buf))
	scanner.Split(bufio.ScanBytes)

	var x, y int
	var width int
	var height int
	var locs []image.Point
	for scanner.Scan() {
		b := scanner.Bytes()[0]
		switch b {
		case '.':
			x++
		case '#':
			locs = append(locs, image.Pt(x, y))
			x++
		case '\n':
			y++
			width = x
			x = 0
		default:
			panic("unexpected byte: " + string(b))
		}
	}
	height = y

	expanded := ExpandSpace(locs, width, height, (amount - 1))

	// TODO: Combinations is returning 10 combinations for 9 locations. :sads:
	pairs := lib.Combinations(expanded, 2)
	fmt.Println("found", len(pairs), "2-combinations")

	sum := 0
	for _, pair := range pairs {
		length := ShortestPath(pair[0], pair[1])
		//fmt.Printf("%d, %d -> %d, %d = %d\n", pair[0].X, pair[0].Y, pair[1].X, pair[1].Y, length)
		sum += length
	}

	return sum
}

func ExpandSpace(locs []image.Point, width, height int, amount int) []image.Point {
	anyX := func(x int) bool {
		for _, l := range locs {
			if l.X == x {
				return true
			}
		}
		return false
	}

	anyY := func(y int) bool {
		for _, l := range locs {
			if l.Y == y {
				return true
			}
		}
		return false
	}

	var res = make([]image.Point, len(locs))
	copy(res, locs)

	for x := 0; x < width; x++ {
		if !anyX(x) {
			fmt.Printf("x=%d is empty, shifting everything >%d to the right\n", x, x)
			for idx := range res {
				if locs[idx].X > x {
					res[idx].X += amount
				}
			}
		}
	}

	for y := 0; y < height; y++ {
		if !anyY(y) {
			fmt.Printf("y=%d is empty, shifting everything >%d down\n", y, y)
			for idx := range res {
				if locs[idx].Y > y {
					res[idx].Y += amount
				}
			}
		}
	}

	return res
}

func ShortestPath(left, right image.Point) int {
	return AbsInt(right.X-left.X) + AbsInt(right.Y-left.Y)
}

func AbsInt(i int) int {
	if i < 0 {
		return -1 * i
	} else {
		return i
	}
}

func (s Solution) Part1(input io.Reader) int {
	return s.Solve(input, 2)
}
func (s Solution) Part2(input io.Reader) int {
	return s.Solve(input, 1_000_000)
}
