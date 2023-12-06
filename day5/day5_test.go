package day5

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

const ExampleInput = "seeds: 79 14 55 13\n\nseed-to-soil map:\n50 98 2\n52 50 48\n\nsoil-to-fertilizer map:\n0 15 37\n37 52 2\n39 0 15\n\nfertilizer-to-water map:\n49 53 8\n0 11 42\n42 0 7\n57 7 4\n\nwater-to-light map:\n88 18 7\n18 25 70\n\nlight-to-temperature map:\n45 77 23\n81 45 19\n68 64 13\n\ntemperature-to-humidity map:\n0 69 1\n1 0 69\n\nhumidity-to-location map:\n60 56 37\n56 93 4\n"

func TestSolution_Part1(t *testing.T) {
	input := bytes.NewBufferString(ExampleInput)

	result := Solution{}.Part1(input)

	require.Equal(t, 35, result)
}

func TestSolution_Part2(t *testing.T) {
	input := bytes.NewBufferString(ExampleInput)

	result := Solution{}.Part2(input)

	require.Equal(t, 46, result)
}
