package day13

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

const ExampleInput = "#.##..##.\n..#.##.#.\n##......#\n##......#\n..#.##.#.\n..##..##.\n#.#.##.#.\n\n#...##..#\n#....#..#\n..##..###\n#####.##.\n#####.##.\n..##..###\n#....#..#\n"

func TestSolution_Part1(t *testing.T) {
	input := bytes.NewBufferString(ExampleInput)

	result := Solution{}.Part1(input)

	assert.Equal(t, 405, result)
}

func TestPart1_SpotCheck(t *testing.T) {
	input := bytes.NewBufferString(".#...##..\n..##.#.##\n.#.###...\n###..#.##\n##.#.####\n..#.#..##\n.###...##\n.#...#.##\n#####.#..\n...#..###\n###.##.##\n####...##\n####..###\n###.##.##\n...#..###\n")

	result := Solution{}.Part1(input)

	assert.Equal(t, 8, result)
}

func TestPart1_SpotCheck2(t *testing.T) {
	input := bytes.NewBufferString(".####..#..#..##\n..##...#..#....\n......###...###\n......##.##..#.\n#....#.#.#.#...\n######.###...##\n.####.#.#..##.#\n.......#.#..#..\n.......#.#..#..\n.####.#.#..##.#\n######.###...##\n#....#.#.#.#...\n......##.##.##.\n")

	result := Solution{}.Part1(input)

	assert.Equal(t, 3, result)
}

func TestPart1_SpotCheck3(t *testing.T) {
	input := bytes.NewBufferString(".##......#.####\n....###..###.##\n.....######....\n....#..##.##...\n######....#.###\n#..###.##....#.\n########.#.###.\n#..#####.#..###\n.##..#...####..\n.##....##..####\n#...##.##.###..\n.##...#..######\n#..#.#.#.#.#...\n.##...#..#....#\n.##...#..#....#\n")

	result := Solution{}.Part1(input)

	assert.Equal(t, 1400, result)
}

func TestSolution_Part2(t *testing.T) {
	input := bytes.NewBufferString(ExampleInput)

	result := Solution{}.Part2(input)

	assert.Equal(t, 400, result)
}
