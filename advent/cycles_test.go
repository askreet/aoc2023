package advent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type RepeatingList struct {
	InitialValues   []int
	RepeatingValues []int
}

func (rl *RepeatingList) Get(index int) int {
	if index < len(rl.InitialValues) {
		return rl.InitialValues[index]
	} else {
		return rl.RepeatingValues[(index-len(rl.InitialValues))%len(rl.RepeatingValues)]
	}
}

func TestRepeatingList(t *testing.T) {
	rl := RepeatingList{
		InitialValues:   []int{0, 9, 2},
		RepeatingValues: []int{3, 6, 1, 4},
	}

	var out []int
	for i := 0; i < 20; i++ {
		out = append(out, rl.Get(i))
	}

	assert.Equal(t,
		[]int{0, 9, 2, 3, 6, 1, 4, 3, 6, 1, 4, 3, 6, 1, 4, 3, 6, 1, 4, 3},
		out)
}

type ListIndex struct {
	rl *RepeatingList
	i  int
}

func (l *ListIndex) Eq(other Brentable) bool {
	if other, ok := other.(*ListIndex); ok {
		return l.rl.Get(l.i) == other.rl.Get(other.i)
	} else {
		panic("other is not *ListIndex")
	}
}

func Next(b Brentable) Brentable {
	li := b.(*ListIndex)

	return &ListIndex{li.rl, li.i + 1}
}

func TestBrentCycleDetection(t *testing.T) {
	rl := RepeatingList{
		InitialValues:   []int{0, 9, 2},
		RepeatingValues: []int{3, 6, 1, 4},
	}

	// lam (λ) is the length of the cycle
	// mu  (𝝁) is the earliest index of the cycle start
	lam, mu := Brent(Next, &ListIndex{&rl, 0})

	assert.Equal(t, 3, mu)
	assert.Equal(t, 4, lam)
}
