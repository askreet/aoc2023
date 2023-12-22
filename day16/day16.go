package day16

import (
	"fmt"
	"image"
	"io"

	"github.com/askreet/aoc2023/advent"
)

type Solution struct{}

const (
	// Previous travelled light bits.
	Up    = 0b00000001
	Down  = 0b00000010
	Left  = 0b00000100
	Right = 0b00001000

	// Mirror bits.
	Empty  = 0b0000_0000
	Vert   = 0b0001_0000
	Horiz  = 0b0010_0000
	Slash  = 0b0011_0000
	BSlash = 0b0100_0000
)

type Direction = byte
type Beam struct {
	image.Point
	Direction
}

var DirStrings = map[byte]string{
	Up:    "Up",
	Down:  "Down",
	Left:  "Left",
	Right: "Right",
}

func (b Beam) String() string {
	return fmt.Sprintf("{Beam (%d, %d) %s}", b.X, b.Y, DirStrings[b.Direction])
}

func (b Beam) Continue() Beam {
	switch b.Direction {
	case Up:
		return Beam{Point: image.Pt(b.X, b.Y-1), Direction: b.Direction}
	case Down:
		return Beam{Point: image.Pt(b.X, b.Y+1), Direction: b.Direction}
	case Left:
		return Beam{Point: image.Pt(b.X-1, b.Y), Direction: b.Direction}
	case Right:
		return Beam{Point: image.Pt(b.X+1, b.Y), Direction: b.Direction}
	default:
		panic("unknown direction")
	}
}

func (b Beam) Up() Beam    { return Beam{Point: image.Pt(b.X, b.Y-1), Direction: Up} }
func (b Beam) Down() Beam  { return Beam{Point: image.Pt(b.X, b.Y+1), Direction: Down} }
func (b Beam) Left() Beam  { return Beam{Point: image.Pt(b.X-1, b.Y), Direction: Left} }
func (b Beam) Right() Beam { return Beam{Point: image.Pt(b.X+1, b.Y), Direction: Right} }

func (s Solution) Part1(input io.Reader) int {
	m := PrepareMap(input)

	return CountVisitsFrom(&m, 0, 0, Right)
}

func (s Solution) Part2(input io.Reader) int {
	m := PrepareMap(input)
	maxFound := 0

	for y := 0; y < m.Height; y++ {
		mCopy := m.Copy()

		if result := CountVisitsFrom(mCopy, 0, y, Right); result > maxFound {
			maxFound = result
		}
	}

	for x := 0; x < m.Width; x++ {
		mCopy := m.Copy()

		if result := CountVisitsFrom(mCopy, x, 0, Down); result > maxFound {
			maxFound = result
		}
	}

	for y := 0; y < m.Height; y++ {
		mCopy := m.Copy()

		if result := CountVisitsFrom(mCopy, m.Width-1, y, Left); result > maxFound {
			maxFound = result
		}
	}

	for x := 0; x < m.Width; x++ {
		mCopy := m.Copy()

		if result := CountVisitsFrom(mCopy, x, m.Height-1, Up); result > maxFound {
			maxFound = result
		}
	}

	return maxFound
}

func PrepareMap(input io.Reader) advent.Map {
	data, err := io.ReadAll(input)
	if err != nil {
		panic(err)
	}

	// Convert input to our custom bits to free space for tracking travelled light paths.
	for i := range data {
		switch data[i] {
		case '/':
			data[i] = Slash
		case '|':
			data[i] = Vert
		case '\\':
			data[i] = BSlash
		case '-':
			data[i] = Horiz
		case '.':
			data[i] = Empty
		}
	}

	m := advent.NewMap(data)
	return m
}

func CountVisitsFrom(m *advent.Map, x, y int, d Direction) int {
	frontier := advent.Stack[Beam]{}
	frontier.Push(Beam{Point: image.Pt(x, y), Direction: d})

	for !frontier.IsEmpty() {
		beam := frontier.Pop()
		if !m.InBounds(beam.X, beam.Y) {
			continue
		}

		// Check if the data on the map has already seen a beam in this direction.
		if m.At(beam.X, beam.Y)&beam.Direction == beam.Direction {
			continue
		}

		// Mark the data on the map as having seen a beam in this direction.
		m.Set(beam.X, beam.Y, m.At(beam.X, beam.Y)|beam.Direction)

		switch m.At(beam.X, beam.Y) & 0xF0 {
		case Empty:
			frontier.Push(beam.Continue())
		case Vert:
			switch beam.Direction {
			case Up, Down:
				frontier.Push(beam.Continue())
			case Left, Right:
				frontier.Push(beam.Up())
				frontier.Push(beam.Down())
			}
		case Horiz:
			switch beam.Direction {
			case Up, Down:
				frontier.Push(beam.Left())
				frontier.Push(beam.Right())
			case Left, Right:
				frontier.Push(beam.Continue())
			}
		case Slash:
			switch beam.Direction {
			case Up:
				frontier.Push(beam.Right())
			case Down:
				frontier.Push(beam.Left())
			case Left:
				frontier.Push(beam.Down())
			case Right:
				frontier.Push(beam.Up())
			}
		case BSlash:
			switch beam.Direction {
			case Up:
				frontier.Push(beam.Left())
			case Down:
				frontier.Push(beam.Right())
			case Left:
				frontier.Push(beam.Up())
			case Right:
				frontier.Push(beam.Down())
			}
		default:
			panic("unexpected direction")
		}
	}

	anyBeamVisited := 0
	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			if m.At(x, y)&0x0F > 0 {
				anyBeamVisited++
			}
		}
	}

	return anyBeamVisited
}
