package day6

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
)

type Solution struct{}

type Race struct {
	Time         int
	BestDistance int
}

func (r Race) WinsIfHeldFor(ms int) bool {
	return (r.Time-ms)*ms > r.BestDistance
}

func (r Race) WinningRange() InclusiveRange {
	var winningSpace = InclusiveRange{-1, -1}

	// Find the earliest winning value.
	var searchSpace = InclusiveRange{0, r.Time}
	var v = searchSpace.Midpoint()
	for {
		this, next := r.WinsIfHeldFor(v), r.WinsIfHeldFor(v+1)
		if !this && next {
			winningSpace.Start = v + 1
			break
		} else if !this && !next {
			searchSpace.Start = v + 1
		} else if this {
			searchSpace.End = v - 1
		}

		v = searchSpace.Midpoint()
	}

	// Find latest winning value.
	searchSpace = InclusiveRange{0, r.Time}
	v = searchSpace.Midpoint()
	for {
		this, next := r.WinsIfHeldFor(v), r.WinsIfHeldFor(v+1)
		if this && !next {
			winningSpace.End = v
			break
		} else if this && next {
			searchSpace.Start = v + 1
		} else if !this {
			searchSpace.End = v - 1
		}

		v = searchSpace.Midpoint()
	}

	return winningSpace
}

func Parse(in io.Reader) []Race {
	var races []Race
	allInput, err := io.ReadAll(in)
	if err != nil {
		panic(err)
	}

	sep := bytes.IndexByte(allInput, '\n')
	line1 := bytes.NewReader(allInput[5:sep])    // Skip 'Time:'
	line2 := bytes.NewReader(allInput[sep+1+9:]) // Skip 'Distance:'

	for {
		var r Race
		n, err := fmt.Fscan(line1, &r.Time)
		if n == 0 {
			break
		} else if err != nil {
			panic(err)
		}
		_, err = fmt.Fscan(line2, &r.BestDistance)
		if err != nil {
			panic(err)
		}
		races = append(races, r)
	}

	return races
}

func Parse2(in io.Reader) Race {
	allInput, err := io.ReadAll(in)
	if err != nil {
		panic(err)
	}

	sep := bytes.IndexByte(allInput, '\n')
	line1 := bufio.NewScanner(bytes.NewReader(allInput[5:sep])) // Skip 'Time:'
	line1.Split(bufio.ScanWords)

	var timeStr = ""
	for line1.Scan() {
		timeStr += line1.Text()
	}

	line2 := bufio.NewScanner(bytes.NewReader(allInput[sep+1+9:])) // Skip 'Distance:'
	line2.Split(bufio.ScanWords)

	var distanceStr = ""
	for line2.Scan() {
		distanceStr += line2.Text()
	}

	time, err := strconv.ParseInt(timeStr, 10, 64)
	if err != nil {
		panic(err)
	}

	distance, err := strconv.ParseInt(distanceStr, 10, 64)
	if err != nil {
		panic(err)
	}

	return Race{
		Time:         int(time),
		BestDistance: int(distance),
	}
}

type InclusiveRange struct {
	Start int
	End   int
}

func (r InclusiveRange) Len() int {
	return r.End - r.Start + 1
}

func (r InclusiveRange) Midpoint() int {
	return r.Start + (r.Len() / 2)
}

func (s Solution) Part1(input io.Reader) int {
	races := Parse(input)

	total := 1
	for _, r := range races {
		winners := r.WinningRange()
		total *= winners.Len()
	}

	return total
}

func (s Solution) Part2(input io.Reader) int {
	race := Parse2(input)

	return race.WinningRange().Len()
}
