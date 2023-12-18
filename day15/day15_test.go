package day15

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	assert.Equal(t, 52, Hash("HASH"))
}

func TestSolution_Part1(t *testing.T) {
	in := bytes.NewBufferString("rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7")

	result := Solution{}.Part1(in)

	assert.Equal(t, 1320, result)
}

func TestSolution_Part2(t *testing.T) {
	in := bytes.NewBufferString("rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7")

	result := Solution{}.Part2(in)

	assert.Equal(t, 145, result)
}
