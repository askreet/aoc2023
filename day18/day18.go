package day18

import (
	"cmp"
	"fmt"
	"image"
	"io"
	"math"
	"slices"
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

type Polygon struct {
	Lines []Line

	topLeft     image.Point
	bottomRight image.Point
}

func NewPolygon() Polygon {
	return Polygon{
		topLeft:     image.Pt(math.MaxInt, math.MaxInt),
		bottomRight: image.Pt(math.MinInt, math.MinInt),
	}
}

func (p *Polygon) Add(v Line) {
	p.Lines = append(p.Lines, v)

	// Because the polygon is tested for completeness, we only need to consider one point in the vertex to have
	// considered all points by the end.
	switch {
	case v.End.X < p.topLeft.X:
		p.topLeft.X = v.End.X
	case v.End.Y < p.topLeft.Y:
		p.topLeft.Y = v.End.Y
	case v.End.X > p.bottomRight.X:
		p.bottomRight.X = v.End.X
	case v.End.Y > p.bottomRight.Y:
		p.bottomRight.Y = v.End.Y
	}
}

func (p *Polygon) IsComplete() bool {
	start := p.Lines[0].Start
	end := p.Lines[len(p.Lines)-1].End

	return start == end
}

func (p *Polygon) Area() int {
	sum := 0
	for y := p.topLeft.Y; y <= p.bottomRight.Y; y++ {
		lines := p.linesIntersectingY(y)
		slices.SortFunc(lines, func(a, b Line) int { return cmp.Compare(a.MinX(), b.MinX()) })

		//transitions, start := 0, 0
		var transitionPoints advent.SparseRangeSet
		var nonTransitionLines advent.SparseRangeSet
		//inside := func() bool { return transitions%2 == 1 }
		for _, intersectingLine := range lines {
			// TODO: Skip vertical lines which overlap a horizontal transition line.
			if p.isTransitionLine(&intersectingLine) {
				transitionPoints.Insert(advent.IRange(intersectingLine.MinX(), intersectingLine.MaxX()))
				//if inside() {
				//	ranges.Insert(advent.IRange(start, intersectingLine.MaxX()))
				//	transitions++
				//} else {
				//	start = intersectingLine.MinX()
				//	transitions++
				//}
			} else {
				nonTransitionLines.Insert(advent.IRange(intersectingLine.MinX(), intersectingLine.MaxX()))
				//if inside() {
				//	// Will be captured by surrounding transitions, nothing to do.
				//} else {
				//	ranges.Insert(advent.IRange(intersectingLine.MinX(), intersectingLine.MaxX()))
				//}
			}
		}
		var coveredAreas advent.SparseRangeSet
		if len(transitionPoints.Ranges)%2 != 0 {
			panic("uneven transition points")
		}
		for i := 0; i < len(transitionPoints.Ranges); i += 2 {
			coveredAreas.Insert(advent.IRange(transitionPoints.Ranges[i].Start, transitionPoints.Ranges[i+1].End))
		}
		for _, r := range nonTransitionLines.Ranges {
			coveredAreas.Insert(r)
		}

		sum += coveredAreas.Total()
	}

	return sum
}

func (p *Polygon) linesIntersectingY(y int) []Line {
	var lines []Line

	for _, l := range p.Lines {
		if l.MinY() <= y && l.MaxY() >= y {
			lines = append(lines, l)
		}
	}

	return lines
}

const (
	Up   = 0
	Down = 1
)

func (p *Polygon) isTransitionLine(l *Line) bool {
	if l.IsVertical() {
		return true
	}

	leftAttachedLine := p.mustFindLineBy(
		func(other Line) bool { return other.IsVertical() && other.HasPoint(l.LeftmostPoint()) })
	rightAttachedLine := p.mustFindLineBy(
		func(other Line) bool { return other.IsVertical() && other.HasPoint(l.RightmostPoint()) })

	leftFacing := Down
	if leftAttachedLine.OtherEnd(l.LeftmostPoint()).Y < l.LeftmostPoint().Y {
		leftFacing = Up
	}

	rightFacing := Down
	if rightAttachedLine.OtherEnd(l.RightmostPoint()).Y < l.RightmostPoint().Y {
		rightFacing = Up
	}

	return leftFacing != rightFacing
}

func (p *Polygon) mustFindLineBy(f func(other Line) bool) *Line {
	idx := slices.IndexFunc(p.Lines, f)
	if idx == -1 {
		panic("line not found in mustFindLineBy")
	}

	return &p.Lines[idx]
}

/*
 XXX...XXX
 X.X...X.X
 X.XXXXX.X..XXXX
 X.......XXXX..X

 XXX
 X.X.....XXXX
 X.XXXXX.X..XXXX
 X.....XXX.....X
*/

type Line struct {
	Start image.Point
	End   image.Point
}

func (l *Line) IsVertical() bool {
	return l.Start.X == l.End.X
}

func (l *Line) IsHorizontal() bool {
	return l.Start.Y == l.End.Y
}

func (l *Line) MinX() int {
	if l.Start.X <= l.End.X {
		return l.Start.X
	} else {
		return l.End.X
	}
}

func (l *Line) MaxX() int {
	if l.Start.X >= l.End.X {
		return l.Start.X
	} else {
		return l.End.X
	}
}

func (l *Line) MinY() int {
	if l.Start.Y <= l.End.Y {
		return l.Start.Y
	} else {
		return l.End.Y
	}
}

func (l *Line) MaxY() int {
	if l.Start.Y >= l.End.Y {
		return l.Start.Y
	} else {
		return l.End.Y
	}
}

func (l *Line) Width() int {
	return l.MaxX() - l.MinX() + 1
}

func (l *Line) LeftmostPoint() image.Point {
	if l.Start.X <= l.End.X {
		return l.Start
	} else {
		return l.End
	}
}

func (l *Line) RightmostPoint() image.Point {
	if l.Start.X >= l.End.X {
		return l.Start
	} else {
		return l.End
	}
}

func (l *Line) HasPoint(point image.Point) bool {
	return l.Start == point || l.End == point
}

func (l *Line) OtherEnd(from image.Point) image.Point {
	if !l.HasPoint(from) {
		panic("OtherEnd called with invalid point")
	}
	if l.Start == from {
		return l.End
	} else {
		return l.Start
	}
}

func (s Solution) Solve(input io.Reader, shouldFix bool) int {
	iter := NewInstIter(input)

	var poly = NewPolygon()
	current := image.Pt(0, 0)
	for iter.Scan() {
		inst := iter.Inst()
		if shouldFix {
			inst = inst.Fix()
		}

		var end image.Point
		switch inst.Dir {
		case "U":
			end = image.Pt(current.X, current.Y-inst.N)

		case "D":
			end = image.Pt(current.X, current.Y+inst.N)

		case "L":
			end = image.Pt(current.X-inst.N, current.Y)

		case "R":
			end = image.Pt(current.X+inst.N, current.Y)

		default:
			panic("unknown dir")
		}
		poly.Add(Line{
			Start: current,
			End:   end,
		})
		current = end
	}

	if !poly.IsComplete() {
		panic("polygon is incomplete")
	}

	return poly.Area()
}

func (s Solution) Part1(input io.Reader) int {
	return s.Solve(input, false)
}

func (s Solution) Part2(input io.Reader) int {
	return s.Solve(input, true)
}
