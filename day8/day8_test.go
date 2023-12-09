package day8

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

const TestInput2 = "LLR\n\nAAA = (BBB, BBB)\nBBB = (AAA, ZZZ)\nZZZ = (ZZZ, ZZZ)\n"
const Part2Example = "LR\n\n11A = (11B, XXX)\n11B = (XXX, 11Z)\n11Z = (11B, XXX)\n22A = (22B, XXX)\n22B = (22C, 22C)\n22C = (22Z, 22Z)\n22Z = (22B, 22B)\nXXX = (XXX, XXX)\n"

func TestSolution_Part1(t *testing.T) {
	input := bytes.NewBufferString(TestInput2)

	assert.Equal(t, 6, Solution{}.Part1(input))
}

func TestSolution_Part2(t *testing.T) {
	input := bytes.NewBufferString(Part2Example)

	assert.Equal(t, 6, Solution{}.Part2(input))
}

func TestLCM(t *testing.T) {
	assert.Equal(t, 48, LCM([]int{12, 16, 24}))
}
