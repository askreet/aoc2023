package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinations(t *testing.T) {
	in := []string{"R", "G", "Y", "B", "I"}

	out := [][]string{
		{"R", "G", "Y"},
		{"R", "G", "B"},
		{"R", "G", "I"},
		{"R", "Y", "B"},
		{"R", "Y", "I"},
		{"R", "B", "I"},
		{"G", "Y", "B"},
		{"G", "Y", "I"},
		{"G", "B", "I"},
		{"Y", "B", "I"},
	}

	assert.Equal(t, out, Combinations(in, 3))
}
