package advent

import "io"

type Interface interface {
	Part1(input io.Reader) int
	Part2(input io.Reader) int
}
