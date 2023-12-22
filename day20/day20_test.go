package day20

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

const FirstExample = "broadcaster -> a, b, c\n%a -> b\n%b -> c\n%c -> inv\n&inv -> a"
const SecondExample = "broadcaster -> a\n%a -> inv, con\n&inv -> b\n%b -> con\n&con -> output"

func TestCreateWorld(t *testing.T) {
	w := CreateWorld(bytes.NewBufferString(FirstExample))

	assert.Equal(t,
		map[string]Module{
			"broadcaster": &Broadcaster{To: []string{"a", "b", "c"}},
			"a":           &FlipFlop{To: []string{"b"}},
			"b":           &FlipFlop{To: []string{"c"}},
			"c":           &FlipFlop{To: []string{"inv"}},
			"inv":         &Conjunction{To: []string{"a"}},
		},
		w.Modules)
}

func TestSolution_Part1_FirstExample(t *testing.T) {
	input := bytes.NewBufferString(FirstExample)

	result := Solution{}.Part1(input)

	assert.Equal(t, 32000000, result)
}

func TestSolution_Part1_SecondExample(t *testing.T) {
	input := bytes.NewBufferString(SecondExample)

	result := Solution{}.Part1(input)

	assert.Equal(t, 11687500, result)
}
