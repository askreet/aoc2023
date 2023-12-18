package day14

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/askreet/aoc2023/advent"
)

const ExampleInput = "O....#....\nO.OO#....#\n.....##...\nOO.#O....O\n.O.....O#.\nO.#..O.#.#\n..O..#O..O\n.......O..\n#....###..\n#OO..#....\n"

func TestSolution_Part1(t *testing.T) {
	in := bytes.NewBufferString(ExampleInput)

	result := Solution{}.Part1(in)

	assert.Equal(t, 136, result)
}

func TestSolution_Part2(t *testing.T) {
	in := bytes.NewBufferString(ExampleInput)

	result := Solution{}.Part2(in)

	assert.Equal(t, 64, result)
}

func TestSlideNorth_BlockedByHash(t *testing.T) {
	m := advent.NewMap([]byte(".\n#\nO\nO\n"))
	SlideNorth(&m)
	assert.Equal(t, ".\n#\nO\nO\n", string(m.Bytes))
}

func TestSlideNorth(t *testing.T) {
	bytes := []byte(`
O..
.O.
..O`)
	m := advent.NewMap(bytes[1:])

	SlideNorth(&m)

	// advent.Map inserts a trailing newline when one is missing.
	expected := []byte(`
OOO
...
...
`)

	assert.Equal(t, expected[1:], m.Bytes)
}

func TestSlideEast(t *testing.T) {
	bytes := []byte(`
O..
.O.
..O`)
	m := advent.NewMap(bytes[1:])

	Slide(&m, East)

	// advent.Map inserts a trailing newline when one is missing.
	expected := []byte(`
..O
..O
..O
`)

	assert.Equal(t, expected[1:], m.Bytes)
}
