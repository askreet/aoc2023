package day10

import (
	"bytes"
	"io"
)

type NodeId int

type Map struct {
	bytes []byte
	width int
}

func (m *Map) ReadFrom(input io.Reader) {
	var err error
	m.bytes, err = io.ReadAll(input)
	if err != nil {
		panic(err)
	}

	m.width = bytes.IndexByte(m.bytes, '\n') + 1
}

func (m *Map) Find(b byte) NodeId {
	return NodeId(bytes.IndexByte(m.bytes, b))
}

func (m *Map) Up(pos NodeId) (NodeId, bool) {
	newpos := int(pos) - m.width

	if newpos <= 0 {
		return 0, false
	}

	return NodeId(newpos), true
}

func (m *Map) Down(pos NodeId) (NodeId, bool) {
	newpos := int(pos) + m.width

	if newpos > len(m.bytes) {
		return 0, false
	}

	return NodeId(newpos), true
}

func (m *Map) Left(pos NodeId) (NodeId, bool) {
	if pos == 0 {
		return 0, false
	}

	newpos := pos - 1
	b := m.bytes[newpos]
	if b == '\n' {
		return 0, false
	}

	return newpos, true
}

func (m *Map) Right(pos NodeId) (NodeId, bool) {
	if int(pos) == len(m.bytes)-1 {
		return 0, false
	}

	newpos := pos + 1
	b := m.bytes[newpos]
	if b == '\n' {
		return 0, false
	}

	return newpos, true
}

func (m *Map) At(pos NodeId) byte {
	return m.bytes[pos]
}

func (m *Map) ConnsAt(n NodeId) []NodeId {
	nw := NodeId(m.width)

	// This function assumes pipes are properly connected within the main loop.
	switch m.At(n) {
	case 'J':
		return []NodeId{n - 1, n - nw}
	case 'F':
		return []NodeId{n + 1, n + nw}
	case 'L':
		return []NodeId{n + 1, n - nw}
	case '7':
		return []NodeId{n - 1, n + nw}
	case '|':
		return []NodeId{n - nw, n + nw}
	case '-':
		return []NodeId{n - 1, n + 1}
	default:
		panic("unexpected byte")
	}
}
