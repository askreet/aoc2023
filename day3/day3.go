package day3

import (
	"io"
	"slices"
)

type Solution struct{}

func (_ Solution) Part1(input io.Reader) int {
	s := NewSchematic(input)
	s.ScanNumbers()

	total := 0
	for _, candidate := range s.PartNumbers {
		isPartNumber := false
		for x := candidate.startX; x <= candidate.endX; x++ {
			if s.cells[s.indexOf(x, candidate.y)].adjSymbol {
				isPartNumber = true
			}
		}

		if isPartNumber {
			total += candidate.value
		}

	}

	return total
}

func (_ Solution) Part2(input io.Reader) int {
	s := NewSchematic(input)
	s.ScanNumbers()

	sum := 0
	for y := 0; y <= s.MaxY; y++ {
		for x := 0; x <= s.MaxX; x++ {
			cell := s.cells[s.indexOf(x, y)]
			if cell.value == '*' {
				var adjPartNumbers []PartNumber
				for _, num := range s.PartNumbers {
					addIfTouches := func(x, y int) {
						if num.Includes(x, y) && !slices.Contains(adjPartNumbers, num) {
							adjPartNumbers = append(adjPartNumbers, num)
						}
					}

					addIfTouches(x-1, y-1)
					addIfTouches(x-1, y)
					addIfTouches(x-1, y+1)

					addIfTouches(x, y-1)

					addIfTouches(x, y+1)

					addIfTouches(x+1, y-1)
					addIfTouches(x+1, y)
					addIfTouches(x+1, y+1)
				}
				if len(adjPartNumbers) == 2 {
					sum += adjPartNumbers[0].value * adjPartNumbers[1].value
				}
			}
		}
	}
	return sum
}

type Schematic struct {
	MaxX, MaxY  int
	cells       []Cell
	PartNumbers []PartNumber
}

type PartNumber struct {
	value  int
	y      int
	startX int
	endX   int
}

func (n PartNumber) Includes(x int, y int) bool {
	return y == n.y && (x >= n.startX && x <= n.endX)
}

type Cell struct {
	value     byte
	adjSymbol bool
}

func NewSchematic(in io.Reader) *Schematic {
	data, err := io.ReadAll(in)
	if err != nil {
		panic(err)
	}

	sch := &Schematic{}

	x, y := 0, 0
	for _, byt := range data {
		if byt == '\n' {
			sch.MaxX = x - 1
			x = 0
			y++
		} else {
			sch.cells = append(sch.cells, Cell{
				value: byt,
			})
			x++
		}
	}

	// If we hit a trailing newline, correct y to be our max value.
	if x == 0 {
		sch.MaxY = y - 1
	} else {
		sch.MaxY = y
	}

	return sch
}

func (s *Schematic) ScanNumbers() {
	for y := 0; y <= s.MaxY; y++ {
		var num PartNumber

		for x := 0; x <= s.MaxX; x++ {
			cell := s.cells[s.indexOf(x, y)]

			if cell.value >= '0' && cell.value <= '9' {
				if num.value == 0 {
					num.startX = x
				}
				num.value = (num.value * 10) + (int(cell.value) - 48)
			} else if num.value > 0 {
				num.endX = x - 1
				num.y = y
				s.PartNumbers = append(s.PartNumbers, num)
				num.value = 0
			}

			if (cell.value < '0' || cell.value > '9') && cell.value != '.' {
				s.flagSymbol(x, y)
			}
		}

		if num.value > 0 {
			num.endX = s.MaxX
			num.y = y
			s.PartNumbers = append(s.PartNumbers, num)
			num.value = 0
		}
	}

}

func (s *Schematic) flagSymbol(x, y int) {
	flagAt := func(x, y int) {
		if x >= 0 && x <= s.MaxX && y >= 0 && y <= s.MaxY {
			s.cells[s.indexOf(x, y)].adjSymbol = true
		}
	}

	flagAt(x-1, y-1)
	flagAt(x-1, y)
	flagAt(x-1, y+1)

	flagAt(x, y-1)

	flagAt(x, y+1)

	flagAt(x+1, y-1)
	flagAt(x+1, y)
	flagAt(x+1, y+1)
}

func (s *Schematic) indexOf(x, y int) int {
	return x + (y * (s.MaxX + 1))
}
