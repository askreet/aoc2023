package day16

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

const Example = ".|...\\....\n|.-.\\.....\n.....|-...\n........|.\n..........\n.........\\\n..../.\\\\..\n.-.-/..|..\n.|....-|.\\\n..//.|....\n"

func TestSolution_Part1(t *testing.T) {
	input := bytes.NewBufferString(Example)

	result := Solution{}.Part1(input)

	assert.Equal(t, 46, result)
}

func TestSolution_Part2(t *testing.T) {
	input := bytes.NewBufferString(Example)

	result := Solution{}.Part2(input)

	assert.Equal(t, 51, result)
}
