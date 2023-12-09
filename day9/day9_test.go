package day9

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextNumber(t *testing.T) {
	assert.Equal(t, 0, NextNumber(0, 0, 0, 0))
	assert.Equal(t, 3, NextNumber(3, 3, 3, 3, 3))
	assert.Equal(t, 18, NextNumber(0, 3, 6, 9, 12, 15))
	assert.Equal(t, 68, NextNumber(10, 13, 16, 21, 30, 45))
}

const Part1Sample = "0 3 6 9 12 15\n1 3 6 10 15 21\n10 13 16 21 30 45\n"

func TestSolution_Part1(t *testing.T) {
	result := Solution{}.Part1(bytes.NewBufferString(Part1Sample))

	assert.Equal(t, 114, result)
}

func Test_SpotCheckPart1(t *testing.T) {
	result := NextNumber(11, 12, 8, -7, -28, -26, 65, 361, 1041, 2360, 4662, 8393, 14114, 22514, 34423, 50825, 72871, 101892, 139412, 187161, 247088)

	assert.Equal(t, 321374, result)
}

func TestSolution_Part2(t *testing.T) {
	result := Solution{}.Part2(bytes.NewBufferString(Part1Sample))

	assert.Equal(t, 2, result)
}
