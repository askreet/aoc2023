package day6

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

const ExampleInput = "Time:      7  15   30\nDistance:  9  40  200\n"

func TestSolution_Part1_Example(t *testing.T) {
	input := bytes.NewBufferString(ExampleInput)

	result := Solution{}.Part1(input)

	assert.Equal(t, 288, result)
}

func TestSolution_Part2_Example(t *testing.T) {
	input := bytes.NewBufferString(ExampleInput)

	result := Solution{}.Part2(input)

	assert.Equal(t, 71503, result)
}

func TestRace_WinningRange(t *testing.T) {
	race := Race{
		Time:         30,
		BestDistance: 200,
	}

	assert.Equal(t, InclusiveRange{11, 19}, race.WinningRange())
}
