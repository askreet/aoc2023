package day15

import (
	"bufio"
	"bytes"
	"io"
	"slices"
	"strconv"

	"github.com/askreet/aoc2023/advent"
)

type Solution struct{}

func Hash(in string) int {
	c := 0
	for i := 0; i < len(in); i++ {
		c += int(in[i])
		c *= 17
		c %= 256
	}
	return c
}

func (s Solution) Part1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(advent.ScanCommaSeperated)

	sum := 0
	for scanner.Scan() {
		sum += Hash(scanner.Text())
	}

	return sum
}

type Lens struct {
	Label       string
	FocalLength int
}

func (s Solution) Part2(input io.Reader) int {
	var boxes [256][]Lens

	scanner := bufio.NewScanner(input)
	scanner.Split(advent.ScanCommaSeperated)

	for scanner.Scan() {
		inst := scanner.Bytes()

		if idx := bytes.IndexByte(inst, '='); idx != -1 {
			key := string(inst[0:idx])
			boxId := Hash(key)
			value, err := strconv.ParseInt(string(inst[idx+1:]), 10, 32)
			if err != nil {
				panic("int conv err: " + err.Error())
			}
			currentIdx := slices.IndexFunc(boxes[boxId], LensWithLabel(key))
			if currentIdx == -1 {
				boxes[boxId] = append(boxes[boxId], Lens{key, int(value)})
			} else {
				boxes[boxId][currentIdx].FocalLength = int(value)
			}

		} else if idx := bytes.IndexByte(inst, '-'); idx != -1 {
			key := string(inst[0:idx])
			boxId := Hash(key)
			boxes[boxId] = slices.DeleteFunc(boxes[boxId], LensWithLabel(key))

		} else {
			panic("unhandled instruction: " + string(inst))
		}
	}

	sum := 0
	for boxId := 0; boxId < len(boxes); boxId++ {
		for lensSlot := 0; lensSlot < len(boxes[boxId]); lensSlot++ {
			sum += (boxId + 1) * (lensSlot + 1) * boxes[boxId][lensSlot].FocalLength
		}
	}
	return sum
}

func LensWithLabel(key string) func(lens Lens) bool {
	return func(lens Lens) bool {
		return lens.Label == key
	}
}
