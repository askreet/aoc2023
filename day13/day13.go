package day13

import (
	"bufio"
	"fmt"
	"io"

	"github.com/askreet/aoc2023/advent"
)

type Solution struct{}
type Map = advent.Map

func (s Solution) Solve(input io.Reader, RefFinder func(Map) Reflection) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(advent.ScanSections)

	n := 1
	sum := 0
	for scanner.Scan() {
		m := advent.NewMap(scanner.Bytes())

		ref := RefFinder(m)
		DisplayMap(n, m, ref)
		sum += ref.Value()

		n++
	}

	return sum
}

func DisplayMap(id int, m Map, r Reflection) {
	fmt.Printf("== Map %d ==\n", id)
	if r.isHorizontal {
		y := 0
		for idx := 0; idx < len(m.Bytes); idx++ {
			if m.Bytes[idx] == '\n' {
				if y == r.leftOrTop {
					fmt.Print("v")
				} else if y == r.leftOrTop+1 {
					fmt.Print("^")
				}
				y++
			}
			fmt.Print(string(m.Bytes[idx]))
		}
	} else {
		for col := 0; col < m.Width; col++ {
			switch {
			case col == r.leftOrTop:
				fmt.Print(">")
			case col == r.leftOrTop+1:
				fmt.Print("<")
			default:
				fmt.Print(" ")
			}
			if col == r.leftOrTop {
			}
		}
		fmt.Println()
		fmt.Println(string(m.Bytes))
	}
	fmt.Println("value:", r.Value())
	fmt.Println()
}

type Reflection struct {
	isHorizontal bool
	leftOrTop    int
}

func (r Reflection) Value() int {
	if r.isHorizontal {
		return (r.leftOrTop + 1) * 100
	} else {
		return r.leftOrTop + 1
	}
}

func FindReflection(m Map) Reflection {
	// Look for horizontal reflection at each y, except the last row.
	for y := 0; y < m.Height-1; y++ {
		// How far from the mirror are we checking?
		for yDelta := 0; y-yDelta >= 0 && y+yDelta+1 < m.Height; yDelta++ {
			for x := 0; x < m.Width; x++ {
				left := m.At(x, y-yDelta)
				right := m.At(x, y+yDelta+1)
				if left != right {
					goto NextRow
				}
			}
		}

		return Reflection{isHorizontal: true, leftOrTop: y}
	NextRow:
	}

	// Look for vertical reflection at each x, except the last column.
	for x := 0; x < m.Width-1; x++ {
		for xDelta := 0; x-xDelta >= 0 && x+xDelta+1 < m.Width; xDelta++ {
			for y := 0; y < m.Height; y++ {
				left := m.At(x-xDelta, y)
				right := m.At(x+xDelta+1, y)
				if left != right {
					goto NextColumn
				}
			}
		}

		return Reflection{isHorizontal: false, leftOrTop: x}
	NextColumn:
	}

	panic("did not find reflection")
}

func FindDirtyReflection(m Map) Reflection {
	// How many differences exist in the reflection at this row or column? If we hit 2, we can give up, since we're looking
	// for a reflection with exactly one defect.
	var nDefects int = 0

	// Look for horizontal reflection at each y, except the last row.
	for y := 0; y < m.Height-1; y++ {
		nDefects = 0

		// How far from the mirror are we checking?
		for yDelta := 0; y-yDelta >= 0 && y+yDelta+1 < m.Height; yDelta++ {
			for x := 0; x < m.Width; x++ {
				top := m.At(x, y-yDelta)
				bottom := m.At(x, y+yDelta+1)

				if top != bottom {
					nDefects++
					if nDefects >= 2 {
						goto NextRow
					}
				}
			}
		}

		if nDefects == 1 {
			return Reflection{isHorizontal: true, leftOrTop: y}
		}
	NextRow:
	}

	// Look for vertical reflection at each x, except the last column.
	for x := 0; x < m.Width-1; x++ {
		nDefects = 0

		for xDelta := 0; x-xDelta >= 0 && x+xDelta+1 < m.Width; xDelta++ {
			for y := 0; y < m.Height; y++ {
				left := m.At(x-xDelta, y)
				right := m.At(x+xDelta+1, y)
				if left != right {
					nDefects++
					if nDefects >= 2 {
						goto NextColumn
					}
				}
			}
		}

		if nDefects == 1 {
			return Reflection{isHorizontal: false, leftOrTop: x}
		}
	NextColumn:
	}

	panic("did not find reflection")
}

func (s Solution) Part1(input io.Reader) int {
	return s.Solve(input, FindReflection)
}

func (s Solution) Part2(input io.Reader) int {
	return s.Solve(input, FindDirtyReflection)
}
