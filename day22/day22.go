package day22

import (
	"bufio"
	"fmt"
	"io"
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

func (s *Shape) EachLoc(fn func(x, y, z int) bool) {
	for x := min(s.X1, s.X2); x <= max(s.X1, s.X2); x++ {
		for y := min(s.Y1, s.Y2); y <= max(s.Y1, s.Y2); y++ {
			for z := min(s.Z1, s.Z2); z <= max(s.Z1, s.Z2); z++ {
				if !fn(x, y, z) {
					// Early abort.
					return
				}
			}
		}
	}
}

func Parse(reader io.Reader) *Tower {
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

	var tower Tower
	for i := range shapes {
		tower.Add(shapes[i])
	}

	return &tower
}

func (s Solution) Part1(input io.Reader) int {
	tower := Parse(input).Settle()

	deletable := 0

	for i := range tower.shapes {
		this := tower.Clone()
		this.DeleteIndex(i)
		if this.SettleStep() == 0 {
			deletable++
			if this.SettleStep() != 0 {
				fmt.Println("tower", i, "settled differently on 2nd attempt")
			}
		}
	}

	return deletable
}

func (s Solution) Part2(input io.Reader) int {
	tower := Parse(input).Settle()

	sum := 0

	for i := range tower.shapes {
		this := tower.Clone()
		this.DeleteIndex(i)
		sum += this.SettleCount()
	}

	return sum
}
