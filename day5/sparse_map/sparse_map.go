package sparse_map

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

type Map struct {
	Entries []Entry
}

func (m *Map) Map(in int) int {
	for _, entry := range m.Entries {
		end := entry.InStart + entry.Len
		if in >= entry.InStart && in < end {
			return entry.OutStart + (in - entry.InStart)
		}
	}

	return in
}

type Entry struct {
	InStart  int
	OutStart int
	Len      int
}

func Parse(in io.Reader) *Map {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)

	m := &Map{}

	for scanner.Scan() {
		line := bytes.NewBuffer(scanner.Bytes())
		var e Entry
		n, err := fmt.Fscanf(line, "%d %d %d", &e.OutStart, &e.InStart, &e.Len)
		if n == 0 {
			break
		} else if err != nil {
			panic(err)
		}
		m.Entries = append(m.Entries, e)
	}

	return m
}
