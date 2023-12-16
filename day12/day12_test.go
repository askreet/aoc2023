package day12

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const Part1_Example = "???.### 1,1,3\n.??..??...?##. 1,1,3\n?#?#?#?#?#?#?#? 1,3,1,6\n????.#...#... 4,1,1\n????.######..#####. 1,6,5\n?###???????? 3,2,1\n"

func Test_solve_Part1_Examples(t *testing.T) {
	tests := []struct {
		p                Pattern
		need             []int
		expectCount      int
		expectIterations int
	}{
		{Pattern("???.###"), []int{1, 1, 3}, 1, 17},
		{Pattern(".??..??...?##."), []int{1, 1, 3}, 4, 88},
		{Pattern("?#?#?#?#?#?#?#?"), []int{1, 3, 1, 6}, 1, 7},
		{Pattern("????.#...#..."), []int{4, 1, 1}, 1, 12},
		{Pattern("????.######..#####."), []int{1, 6, 5}, 4, 37},
		{Pattern("?###????????"), []int{3, 2, 1}, 10, 71},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("solve(%v, %v) = %v", string(test.p), test.need, test.expectCount), func(t *testing.T) {
			var tracker AlgoTracker
			n := solve(test.p, test.need, &tracker)

			assert.Equal(t, test.expectCount, n)
			assert.Equal(t, test.expectIterations, tracker.Iterations)
		})
	}
}

func Test_solve_Part2_Examples(t *testing.T) {
	tests := []struct {
		p                Pattern
		need             []int
		expectCount      int
		expectIterations int
	}{
		{Pattern("???.###"), []int{1, 1, 3}, 1, 85},
		{Pattern(".??..??...?##."), []int{1, 1, 3}, 16384, 334708},
		{Pattern("?#?#?#?#?#?#?#?"), []int{1, 3, 1, 6}, 1, 35},
		{Pattern("????.#...#..."), []int{4, 1, 1}, 16, 327},
		{Pattern("????.######..#####."), []int{1, 6, 5}, 2500, 28741},
		{Pattern("?###????????"), []int{3, 2, 1}, 506250, 4021063},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("solve(%v, %v) = %v", string(test.p), test.need, test.expectCount), func(t *testing.T) {
			var tracker AlgoTracker
			unfolded := Unfold(Puzzle{test.p, test.need})
			n := solve(unfolded.p, unfolded.need, &tracker)

			assert.Equal(t, test.expectCount, n)
			assert.Equal(t, test.expectIterations, tracker.Iterations)
		})
	}
}
func Test_solve_BaseCases(t *testing.T) {
	assert.Equal(t, 1, solve(Pattern("#"), []int{1}, nil))
	assert.Equal(t, 1, solve(Pattern(""), []int{}, nil))
	assert.Equal(t, 0, solve(Pattern(""), []int{1}, nil))
}

func Test_solve_EdgeCases(t *testing.T) {
	solve(Pattern("#"), []int{2, 1}, nil)
}

func TestSolution_Part1(t *testing.T) {
	in := bytes.NewBufferString(Part1_Example)

	result := Solution{}.Part1(in)

	assert.Equal(t, 21, result)
}

func TestSolution_Part2(t *testing.T) {
	in := bytes.NewBufferString(Part1_Example)

	result := Solution{}.Part2(in)

	assert.Equal(t, 525152, result)
}

func TestSolution_Part2_SlowExample1(t *testing.T) {
	in := bytes.NewBufferString("?#??#???.??.??? 3,1,1,1,1\n")

	result := Solution{}.Part2(in)

	assert.Equal(t, 10403882, result)
}

func TestSolution_Part2_SlowExample2(t *testing.T) {
	in := bytes.NewBufferString("??#????????????? 5,2,2\n")

	result := Solution{}.Part2(in)

	assert.Equal(t, 10403882, result)
}

func Test_solve_EarlyAbortTooManyNeeds(t *testing.T) {
	p := Pattern("#.?.?.?.?")
	needs := []int{1, 1, 1, 1, 1, 1}

	var trk AlgoTracker
	r := solve(p, needs, &trk)

	assert.Equal(t, trk.Iterations, 1)
	assert.Equal(t, 0, r)
}

func Test_solve_NegativeLookahead(t *testing.T) {
	p := Pattern("???.##???")
	needs := []int{5}

	var trk AlgoTracker
	r := solve(p, needs, &trk)

	assert.Less(t, trk.Iterations, 5)
	assert.Equal(t, 1, r)
}
