package day18

import (
	"fmt"
	"image/color"
	"io"
	"strconv"

	"github.com/askreet/aoc2023/advent"
)

type Solution struct{}

type Inst struct {
	Dir   string
	N     int
	Color string
}

var FixDirs = map[byte]string{
	'0': "R",
	'1': "D",
	'2': "L",
	'3': "U",
}

func (i Inst) Fix() *Inst {
	var fixed Inst

	if dir, ok := FixDirs[i.Color[7]]; ok {
		fixed.Dir = dir
	} else {
		panic("unexpected fixed dir bit")
	}

	hexValue, err := strconv.ParseInt(i.Color[2:7], 16, 32)
	if err != nil {
		panic("failed to parse hex value from " + i.Color[2:7])
	}
	fixed.N = int(hexValue)

	return &fixed
}

type InstIter struct {
	r       io.Reader
	hasNext bool
	next    Inst
}

func NewInstIter(r io.Reader) InstIter {
	return InstIter{r: r}
}

func (i *InstIter) Scan() bool {
	n, err := fmt.Fscanf(i.r, "%s %d %s\n", &i.next.Dir, &i.next.N, &i.next.Color)
	if n < 3 || err != nil {
		return false
	}

	return true
}

func (i *InstIter) Inst() *Inst {
	return &i.next
}

func (s Solution) Solve(input io.Reader, shouldFix bool) int {
	iter := NewInstIter(input)

	x, y := 0, 0
	fmt.Println("staritng infmap")
	var m advent.InfMap
	for iter.Scan() {
		inst := iter.Inst()
		if shouldFix {
			inst = inst.Fix()
		}
		switch inst.Dir {
		case "U":
			for i := 0; i < inst.N; i++ {
				y -= 1
				m.Set(x, y, '#')
			}

		case "D":
			for i := 0; i < inst.N; i++ {
				y += 1
				m.Set(x, y, '#')
			}

		case "L":
			for i := 0; i < inst.N; i++ {
				x -= 1
				m.Set(x, y, '#')
			}

		case "R":
			for i := 0; i < inst.N; i++ {
				x += 1
				m.Set(x, y, '#')
			}
		default:
			panic("unknown dir")
		}
	}
	fmt.Println("infmap ready")

	fmt.Println("creating map")
	nm := m.AsMap('.')
	fmt.Println("map ready")
	colors := map[byte]color.Color{
		'#': color.RGBA{R: 0xFF, A: 0xFF},
		'F': color.RGBA{R: 0x99, A: 0xFF},
	}
	nm.SavePNG("before.png", colors)
	for y := 0; y < nm.Height; y++ {
		var lastByte byte = '.'
		c := 0
		isWide := false
		wideFromTop := false
		for x := 0; x < nm.Width; x++ {
			b := nm.At(x, y)
			switch b {
			case '#':
				if lastByte == '.' {
					isWide = false
					c++
				} else if isWide == false {
					isWide = true
					if nm.InBounds(x-1, y-1) && nm.At(x-1, y-1) == '#' {
						wideFromTop = true
					} else if nm.InBounds(x-1, y+1) && nm.At(x-1, y+1) == '#' {
						wideFromTop = false
					} else {
						panic("could not determine wide pipe opening transition")
					}
				}
				lastByte = b
			case '.':
				if lastByte == '#' && isWide {
					// Determine if our twist is closing the loop or not.
					if nm.InBounds(x-1, y-1) && nm.At(x-1, y-1) == '#' {
						// Turned upward, transition again if we came from up.
						if wideFromTop {
							c++
						}
					} else if nm.InBounds(x-1, y+1) && nm.At(x-1, y+1) == '#' {
						// Turned downward, transition again if we came from down.
						if !wideFromTop {
							c++
						}
					} else {
						panic("could not determine wide pipe closing transition")
					}
				}
				if c%2 == 1 {
					nm.Set(x, y, 'F')
				}
				lastByte = b
			default:
				panic("unknown byte")
			}
		}
	}

	nm.SavePNG("after.png", colors)
	return nm.Count('#') + nm.Count('F')
}

func (s Solution) Part1(input io.Reader) int {
	return s.Solve(input, false)
}

func (s Solution) Part2(input io.Reader) int {
	return s.Solve(input, true)
}
