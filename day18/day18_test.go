package day18

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

const ExampleInput = "R 6 (#70c710)\nD 5 (#0dc571)\nL 2 (#5713f0)\nD 2 (#d2c081)\nR 2 (#59c680)\nD 2 (#411b91)\nL 5 (#8ceee2)\nU 2 (#caa173)\nL 1 (#1b58a2)\nU 2 (#caa171)\nR 2 (#7807d2)\nU 3 (#a77fa3)\nL 2 (#015232)\nU 2 (#7a21e3)\n"

func TestSolution_Part1(t *testing.T) {
	input := bytes.NewBufferString(ExampleInput)

	result := Solution{}.Part1(input)

	assert.Equal(t, 62, result)
}

func TestSolution_Part2(t *testing.T) {
	input := bytes.NewBufferString(ExampleInput)

	result := Solution{}.Part2(input)

	assert.Equal(t, 952408144115, result)
}

// XXX.XXX
// X.X.X.X
// X.XXX.X // <== off by one due to sparse list not collapsing
func TestPart1_EmptySpaceRightOfWideLine(t *testing.T) {
	input := bytes.NewBufferString(
		"D 6 (#ffffff)\n" +
			"R 2 (#ffffff)\n" +
			"U 2 (#ffffff)\n" +
			"R 2 (#ffffff)\n" +
			"D 2 (#ffffff)\n" +
			"R 2 (#ffffff)\n" +
			"U 6 (#ffffff)\n" +
			"L 2 (#ffffff)\n" +
			"D 2 (#ffffff)\n" +
			"L 2 (#ffffff)\n" +
			"U 2 (#ffffff)\n" +
			"L 2 (#ffffff)\n",
	)

	result := Solution{}.Part1(input)

	assert.Equal(t, 45, result)
}

func TestInst_Fix(t *testing.T) {
	subject := Inst{
		Dir:   "R",
		N:     6,
		Color: "(#70c710)",
	}

	fixed := subject.Fix()

	assert.Equal(t, "R", fixed.Dir)
	assert.Equal(t, 461937, fixed.N)
}
