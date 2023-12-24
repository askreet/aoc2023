package day22

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution_Part1(t *testing.T) {
	input := bytes.NewBufferString(ExampleInput)

	result := Solution{}.Part1(input)

	assert.Equal(t, 5, result)
}

func TestSolution_Part1_Input(t *testing.T) {
	input, _ := os.ReadFile("../in/day22.txt")

	result := Solution{}.Part1(bytes.NewBuffer(input))

	assert.Equal(t, 517, result)
}

func Benchmark_Parse(b *testing.B) {
	input := bytes.NewBufferString(ExampleInput)

	for i := 0; i < b.N; i++ {
		Parse(input)
	}
}

func BenchmarkTower_SettleTestInput(b *testing.B) {
	input, _ := os.ReadFile("../in/day22.txt")
	for i := 0; i < b.N; i++ {
		tower := Parse(bytes.NewBuffer(input))
		tower.Settle()
	}
}

func TestSolution_Part2(t *testing.T) {
	input := bytes.NewBufferString(ExampleInput)

	result := Solution{}.Part2(input)

	assert.Equal(t, 7, result)
}
