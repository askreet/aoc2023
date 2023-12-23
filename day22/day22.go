package day22

import (
	"bufio"
	"fmt"
	"io"
	"slices"
)

const ExampleInput = "1,0,1~1,2,1\n0,0,2~2,0,2\n0,2,3~2,2,3\n0,0,4~0,2,4\n2,0,5~2,2,5\n0,1,6~2,1,6\n1,1,8~1,1,9\n"

type Solution struct{}

type Shape struct {
	X1, Y1, Z1 int
	X2, Y2, Z2 int
	bit        bool
}

func (s Shape) MinZ() int {
	return min(s.Z1, s.Z2)
}

func (s Shape) Drop() Shape {
	s.Z1--
	s.Z2--
	s.bit = true
	return s
}

func (s Shape) Overlaps(b Shape) bool {
	overlapX := max(s.X1, s.X2) >= min(b.X1, b.X2) &&
		min(s.X1, s.X2) <= max(b.X1, b.X2)
	overlapY := max(s.Y1, s.Y2) >= min(b.Y1, b.Y2) &&
		min(s.Y1, s.Y2) <= max(b.Y1, b.Y2)
	overlapZ := max(s.Z1, s.Z2) >= min(b.Z1, b.Z2) &&
		min(s.Z1, s.Z2) <= max(b.Z1, b.Z2)

	return overlapX && overlapY && overlapZ
}

func (s *Shape) EachLoc(fn func(x, y, z int)) {
	for x := min(s.X1, s.X2); x <= max(s.X1, s.X2); x++ {
		for y := min(s.Y1, s.Y2); y <= max(s.Y1, s.Y2); x++ {
			for z := min(s.Z1, s.Z2); z <= max(s.Z1, s.Z2); x++ {
				fn(x, y, z)
			}
		}
	}
}

type Tower []Shape

// Settle will generate a new tower where all pieces have moved far down as possible without colliding with another piece.
func (t Tower) Settle() Tower {
	var newTower = make(Tower, len(t))
	copy(newTower, t)

	for newTower.SettleStep() > 0 {
	}

	return newTower
}

// Count the number of bricks that will settle, total.
func (t *Tower) SettleCount() int {
	for i := range *t {
		(*t)[i].bit = false
	}

	for {
		if t.SettleStep() == 0 {
			break
		}
	}

	sum := 0

	for i := range *t {
		if (*t)[i].bit {
			sum++
		}
	}

	return sum
}

// SettleStep executes a single settle action by mutating the current Tower. It returns the number of changed shapes.
func (t *Tower) SettleStep() int {
	numMoved := 0

	var lastSuccessfulLocation Shape
	var candidateLocation Shape
	for this := range *t {
		if (*t)[this].MinZ() > 0 {
			lastSuccessfulLocation = (*t)[this]
			for {
				candidateLocation = lastSuccessfulLocation.Drop()
				collides := false
				for other := range *t {
					if this == other {
						continue
					}

					if candidateLocation.Overlaps((*t)[other]) {
						collides = true
						break
					}
				}
				if collides || candidateLocation.MinZ() == 0 {
					(*t)[this] = lastSuccessfulLocation
					break
				} else {
					lastSuccessfulLocation = candidateLocation
					numMoved++
				}
			}
		}
	}

	return numMoved
}

func (t *Tower) Clone() Tower {
	var n Tower = make(Tower, len(*t))
	copy(n, *t)
	return n
}

func Parse(reader io.Reader) Tower {
	var shapes []Shape

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()

		var shape Shape
		n, err := fmt.Sscanf(line, "%d,%d,%d~%d,%d,%d", &shape.X1, &shape.Y1, &shape.Z1,
			&shape.X2, &shape.Y2, &shape.Z2)
		if n != 6 {
			panic("not enough values extracted for " + line)
		}
		if err != nil {
			panic(err)
		}

		shapes = append(shapes, shape)
	}

	return shapes
}

func (s Solution) Part1(input io.Reader) int {
	shapes := Parse(input)
	shapes = shapes.Settle()

	deletable := 0

	for i := range shapes {
		this := shapes.Clone()
		this = slices.Delete(this, i, i+1)
		if this.SettleStep() == 0 {
			deletable++
		}
	}

	return deletable
}

func (s Solution) Part2(input io.Reader) int {
	shapes := Parse(input)
	shapes = shapes.Settle()

	sum := 0

	for i := range shapes {
		this := shapes.Clone()
		this = slices.Delete(this, i, i+1)
		sum += this.SettleCount()
	}

	return sum
}
