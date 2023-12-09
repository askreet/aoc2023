package day8

import (
	"bufio"
	"fmt"
	"io"
	"math"
)

type Solution struct {
}

type LocId string
type Inst int

const (
	Left  Inst = 0
	Right Inst = 1
)

type Element struct {
	Left  LocId
	Right LocId
}

func (e Element) At(inst Inst) LocId {
	if inst == Left {
		return e.Left
	} else if inst == Right {
		return e.Right
	}
	panic("invalid Inst for Element.At()")
}

type Document struct {
	Instructions RingInstructions
	Elements     map[LocId]Element
}

type RingInstructions struct {
	str string
	idx int
}

func (r *RingInstructions) Next() Inst {
	this := r.idx
	r.idx = r.idx + 1
	if r.idx > len(r.str)-1 {
		r.idx = 0
	}

	switch r.str[this] {
	case 'L':
		return Left
	case 'R':
		return Right
	default:
		panic("unexpected instruction in RingInstructions")
	}
}

func Parse(input io.Reader) Document {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	var doc Document
	scanner.Scan()
	doc.Instructions = RingInstructions{str: scanner.Text()}

	// Blank line.
	scanner.Scan()

	doc.Elements = make(map[LocId]Element, 1024)
	for scanner.Scan() {
		line := scanner.Text()
		doc.Elements[LocId(line[0:3])] = Element{
			Left:  LocId(line[7:10]),
			Right: LocId(line[12:15]),
		}
	}

	return doc
}

func (s Solution) Part1(input io.Reader) int {
	doc := Parse(input)

	steps := 0
	location := LocId("AAA")
	for location != "ZZZ" {
		inst := doc.Instructions.Next()

		location = doc.Elements[location].At(inst)
		steps++
	}

	return steps
}

func (s Solution) Part2(input io.Reader) int {
	doc := Parse(input)

	steps := 0
	var locations []LocId
	for locId := range doc.Elements {
		if locId[2] == 'A' {
			locations = append(locations, locId)
		}
	}
	var firstArrived = make([]int, len(locations))

	allArrivedOnce := func() bool {
		for _, v := range firstArrived {
			if v == 0 {
				return false
			}
		}
		return true
	}

	for !allArrivedOnce() {
		inst := doc.Instructions.Next()

		steps++
		for idx := range locations {
			locations[idx] = doc.Elements[locations[idx]].At(inst)
			if locations[idx][2] == 'Z' && firstArrived[idx] == 0 {
				fmt.Println("ghost", idx, "arrives at Z on step", steps)
				firstArrived[idx] = steps
			}
		}
	}

	return LCM(firstArrived)
}

// LCM calculates the least common multiple of a set of numbers. It is not remotely the most efficient method, but it's
// mine and I like it.
func LCM(numbers []int) int {
	smallest := math.MaxInt
	for _, v := range numbers {
		if v < smallest {
			smallest = v
		}
	}

	this := 0
	for {
		this += smallest

		valid := true
		for _, v := range numbers {
			if this < v || this%v != 0 {
				valid = false
			}
		}

		if valid {
			return this
		}
	}
}
