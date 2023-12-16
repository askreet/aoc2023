package day11

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

const ExamplePartOne = "...#......\n.......#..\n#.........\n..........\n......#...\n.#........\n.........#\n..........\n.......#..\n#...#.....\n"

func TestSolution_Part1(t *testing.T) {
	input := bytes.NewBufferString(ExamplePartOne)

	result := Solution{}.Part1(input)

	assert.Equal(t, 374, result)
}

func TestSolution_Part2_Ex1(t *testing.T) {
	input := bytes.NewBufferString(ExamplePartOne)

	result := Solution{}.Solve(input, 10)

	assert.Equal(t, 1030, result)
}

func TestSolution_Part2_Ex2(t *testing.T) {
	input := bytes.NewBufferString(ExamplePartOne)

	result := Solution{}.Solve(input, 100)

	assert.Equal(t, 8410, result)
}
