package day4

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

const ExampleInput = "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\nCard 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\nCard 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\nCard 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\nCard 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\nCard 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11\n"

func TestSolution_Part1_Example(t *testing.T) {
	result := Solution{}.Part1(bytes.NewBufferString(ExampleInput))

	assert.Equal(t, 13, result)
}
func TestSolution_Part2_Example(t *testing.T) {
	result := Solution{}.Part2(bytes.NewBufferString(ExampleInput))

	assert.Equal(t, 30, result)
}
