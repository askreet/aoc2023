package day10

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

const Part1Example = "..F7.\n.FJ|.\nSJ.L7\n|F--J\nLJ...\n"

func TestSolution_Part1(t *testing.T) {
	input := bytes.NewBufferString(Part1Example)

	result := Solution{}.Part1(input)

	assert.Equal(t, 8, result)
}

const Part2Example1 = "...........\n.S-------7.\n.|F-----7|.\n.||.....||.\n.||.....||.\n.|L-7.F-J|.\n.|..|.|..|.\n.L--J.L--J.\n...........\n"

func TestSolution_Part2(t *testing.T) {
	input := bytes.NewBufferString(Part2Example1)

	result := Solution{}.Part2(input)

	assert.Equal(t, 4, result)
}

const Part2Example2 = ".F----7F7F7F7F-7....\n.|F--7||||||||FJ....\n.||.FJ||||||||L7....\nFJL7L7LJLJ||LJ.L-7..\nL--J.L7...LJS7F-7L7.\n....F-J..F7FJ|L7L7L7\n....L7.F7||L7|.L7L7|\n.....|FJLJ|FJ|F7|.LJ\n....FJL-7.||.||||...\n....L---J.LJ.LJLJ...\n"

func TestSolution_Part2_Ex2(t *testing.T) {
	input := bytes.NewBufferString(Part2Example2)

	result := Solution{}.Part2(input)

	assert.Equal(t, 8, result)
}

const Part2Example3 = "FF7FSF7F7F7F7F7F---7\nL|LJ||||||||||||F--J\nFL-7LJLJ||||||LJL-77\nF--JF--7||LJLJ7F7FJ-\nL---JF-JLJ.||-FJLJJ7\n|F|F-JF---7F7-L7L|7|\n|FFJF7L7F-JF7|JL---7\n7-L-JL7||F7|L7F-7F7|\nL.L7LFJ|||||FJL7||LJ\nL7JLJL-JLJLJL--JLJ.L\n"

func TestSolution_Part2_Ex3(t *testing.T) {
	input := bytes.NewBufferString(Part2Example3)

	result := Solution{}.Part2(input)

	assert.Equal(t, 10, result)
}
