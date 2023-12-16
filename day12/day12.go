package day12

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"time"
)

const DebugAlgo = false

type Solution struct{}

type Puzzle struct {
	p    Pattern
	need []int
}

func Parse(in io.Reader) []Puzzle {
	var puzzles []Puzzle

	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanBytes)

	var puzzle Puzzle
	var n int
	for scanner.Scan() {
		b := scanner.Bytes()[0]

		switch b {
		case '#', '?', '.':
			puzzle.p = append(puzzle.p, b)

		case ' ':

		case ',':
			puzzle.need = append(puzzle.need, n)
			n = 0

		case '\n':
			puzzle.need = append(puzzle.need, n)
			n = 0

			puzzles = append(puzzles, puzzle)
			puzzle = Puzzle{}

		default:
			if n == 0 {
				n = int(b & 0b0000_1111)
			} else {
				n *= 10
				n += int(b & 0b0000_1111)
			}
		}
	}

	return puzzles
}

func (s Solution) Part1(input io.Reader) int {
	puzzles := Parse(input)

	sum := 0
	iters := 0
	for _, p := range puzzles {
		var trk AlgoTracker
		c := solve(p.p, p.need, &trk)
		sum += c
		iters += trk.Iterations
	}
	fmt.Println("part 1 iterations =", iters)

	return sum
}

func Unfold(p Puzzle) Puzzle {
	var bs [][]byte
	var needs []int
	for i := 0; i < 5; i++ {
		bs = append(bs, p.p)
		needs = append(needs, p.need...)
	}

	return Puzzle{
		p:    bytes.Join(bs, []byte{'?'}),
		need: needs,
	}
}

func (s Solution) Part2(input io.Reader) int {
	puzzles := Parse(input)

	sum := 0
	iters := 0
	for idx, p := range puzzles {
		start := time.Now()
		unfolded := Unfold(p)
		var trk *AlgoTracker
		if DebugAlgo {
			trk = &AlgoTracker{}
		}
		c := solve(unfolded.p, unfolded.need, trk)
		sum += c
		ms := time.Now().Sub(start).Milliseconds()
		if DebugAlgo {
			fmt.Println("solved puzzle", idx, "in", ms, "ms",
				fmt.Sprintf("%v searched", trk.Iterations))
		}
	}
	fmt.Println("part 2 iterations =", iters)

	return sum
}

type AlgoTracker struct {
	// How many whole solutions did we consider?
	Iterations int
}

type Pattern []byte

func (p Pattern) Copy() Pattern {
	var newPattern Pattern = make([]byte, len(p))
	copy(newPattern, p)
	return newPattern
}

func solve(p Pattern, needs []int, trk *AlgoTracker) int {
	if trk != nil {
		trk.Iterations++
		if DebugAlgo {
			fmt.Println("solve", string(p), needs)
		}
	}

	if len(p) == 0 && len(needs) == 0 {
		return 1
	} else if len(p) == 0 {
		return 0
	}

	switch p[0] {
	case '?':
		// Recurse for each value of ?. Return sum of recursion calls.
		sum := 0
		on := p.Copy()
		on[0] = '#'
		sum += solve(on, needs, trk)

		off := p.Copy()
		off[0] = '.'
		sum += solve(off, needs, trk)

		return sum

	case '#':
		// If there's pattern left to fill out, but no remaining needs, having any # is never correct, so the pattern
		// is not viable.
		if len(needs) == 0 {
			return 0
		}

		// If the sum of the remaining needs plus necessary separators is greater than the length of the pattern,
		// it is also impossible to solve.
		if sum(needs...)+len(needs)-1 > len(p) {
			return 0
		}

		// Check to see if it would be possible to set the next N characters to # to match the current need along
		// with our current #. If not, this isn't a viable pattern.
		lookahead := p[0:needs[0]]
		if bytes.Count(lookahead, []byte{'.'}) > 0 {
			return 0
		}

		// Next, determine if we're at the end of the pattern so we can determine if we met all the needs.
		if len(p) == needs[0] {
			if len(needs) == 1 {
				// No further unmet needs, this is a winner.
				return 1
			} else {
				// We matched this need, but there are further needs and we're out of pattern.
				return 0
			}
		}

		// Since we have further pattern to match against we need to make sure we can _terminate_ this block with a '.'.
		// If not, this is an invalid pattern.
		if p[needs[0]] == '#' {
			return 0
		}

		// And finally since we could put a legal match for the next need, recurse into the remaining pattern with
		// needs[1:]
		nextPattern := p[needs[0]+1:]
		nextNeeds := needs[1:]
		return solve(nextPattern, nextNeeds, trk)

	case '.':
		// If we have no needs left, consider if the rest of the pattern could be interpreted as '.'. If so, that's the
		// only viable matching pattern. If not, there is no viable matching pattern.
		if len(needs) == 0 {
			if bytes.Count(p, []byte{'#'}) > 0 {
				return 0
			} else {
				return 1
			}
		}

		if len(p) < needs[0] {
			return 0
		}

		nextDot := bytes.LastIndexByte(p[0:needs[0]], '.')
		if nextDot != -1 {
			// There's no way to fit the next chunk of #'s in to the next position, so skip to after the last '.' and see
			// if it's a good candidate to match needs[0].
			return solve(p[nextDot+1:], needs, trk)
		} else {
			return solve(p[1:], needs, trk)
		}

	default:
		panic("unexpected char")
	}
}

func sum(nums ...int) int {
	i := 0
	for _, n := range nums {
		i += n
	}
	return i
}
