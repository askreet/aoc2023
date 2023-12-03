package day1

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1_Sample(t *testing.T) {
	input := bytes.NewBufferString("1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet")

	result := part1(input)

	assert.Equal(t, result, 142)
}

func TestPart2_Sample(t *testing.T) {
	input := bytes.NewBufferString("two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen\n")

	result := part2(input)

	assert.Equal(t, result, 281)
}

func TestPart2_EdgeCase1(t *testing.T) {
	// "oneight" needs to match "eight" from the right
	input := bytes.NewBufferString("8kgplfhvtvqpfsblddnineoneighthg\n")

	result := part2(input)

	assert.Equal(t, 88, result)
}
