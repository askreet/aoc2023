package day14

import (
	"bufio"
	"bytes"
	"io"

	"github.com/askreet/aoc2023/advent"
)

type Solution struct{}

func Solve(in io.Reader) int {
	scanner := bufio.NewScanner(in)
	scanner.Split(advent.ScanSections)

	sum := 0
	id := 1
	for scanner.Scan() {
		m := advent.NewMap(scanner.Bytes())
		m.Print(id)

		SlideNorth(&m)
		m.Print(id)

		sum += CalculateLoad(&m)
		id++
	}
	return sum
}

type Translation struct {
	X int
	Y int
}

func (t Translation) ApplyTo(x, y int) (int, int) {
	return x + t.X, y + t.Y
}

func (t Translation) Inverse() Translation {
	return Translation{t.X * -1, t.Y * -1}
}

func (t Translation) VisitOrder(m *advent.Map) *advent.VisitOrder {
	if t.X > 0 || t.Y > 0 {
		return m.OrderInverse()
	} else {
		return m.OrderNormal()
	}
}

func SlideNorth(m *advent.Map) {
	Slide(m, Translation{0, -1})
}

func Slide(m *advent.Map, t Translation) {
	m.VisitAll('O', t.VisitOrder(m), func(x, y int) {
		cX, cY := t.ApplyTo(x, y)

		for {
			if !m.InBounds(cX, cY) {
				goto StopSliding
			}

			b := m.At(cX, cY)
			switch b {
			case '#', 'O':
				goto StopSliding
			case '.':
				// keep sliding
				cX, cY = t.ApplyTo(cX, cY)
			default:
				panic("unexpected byte!")
			}
		}
	StopSliding:
		// Rewind the slide to the previous valid location.
		cX, cY = t.Inverse().ApplyTo(cX, cY)

		if cX != x || cY != y {
			m.Set(cX, cY, 'O')
			m.Set(x, y, '.')
		}
	})
}

func CalculateLoad(m *advent.Map) int {
	load := 0

	m.VisitAll('O', m.OrderNormal(), func(x, y int) {
		additionalLoad := m.Height - y
		load += additionalLoad
	})

	return load
}

func (s Solution) Part1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(advent.ScanSections)

	sum := 0
	id := 1
	for scanner.Scan() {
		m := advent.NewMap(scanner.Bytes())

		SlideNorth(&m)

		sum += CalculateLoad(&m)
		id++
	}
	return sum
}

var (
	North = Translation{0, -1}
	West  = Translation{-1, 0}
	South = Translation{0, 1}
	East  = Translation{1, 0}
)

type MapState struct {
	m *advent.Map
}

func (m MapState) Eq(other advent.Brentable) bool {
	return bytes.Equal(m.m.Bytes, other.(*MapState).m.Bytes)
}

func NextMap(ms advent.Brentable) advent.Brentable {
	if ms, ok := ms.(*MapState); ok {
		var next MapState
		next.m = ms.m.Copy()

		Slide(next.m, North)
		Slide(next.m, West)
		Slide(next.m, South)
		Slide(next.m, East)

		return &next
	} else {
		panic("expect *MapState")
	}
}

func (s Solution) Part2(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(advent.ScanSections)
	scanner.Scan()

	m := advent.NewMap(scanner.Bytes())

	λ, 𝝁 := advent.Brent(NextMap, &MapState{&m})

	state := &MapState{&m}
	iters := 𝝁 + ((1_000_000_000 - 𝝁) % λ)
	for i := 0; i < iters; i++ {
		state = NextMap(state).(*MapState)
	}

	return CalculateLoad(state.m)
}
