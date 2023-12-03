package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

const Example = "467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598.."

func TestPart1_Example(t *testing.T) {
	result := part1(bytes.NewBufferString(Example))

	assert.Equal(t, 4361, result)
}

func TestPart2_Example(t *testing.T) {
	result := part2(bytes.NewBufferString(Example))

	assert.Equal(t, 467835, result)
}
