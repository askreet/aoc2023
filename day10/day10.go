package day10

import (
	"bytes"
	"io"
)

type Solution struct{}

type MapResult struct {
	Map   *Map
	Start NodeId
}

// CreateMap converts the input to a Map, and returns the start node, which has been
// replaced with its actual pipe value allowing algorithms to run against the map.
func CreateMap(input io.Reader) MapResult {
	var map_ Map
	map_.ReadFrom(input)

	startLoc := map_.Find('S')

	hasUp, hasLeft, hasRight, hasDown := false, false, false, false
	// Special case: find adjacent pipes connected to S.
	if up, ok := map_.Up(startLoc); ok {
		b := map_.At(up)
		if b == 'F' || b == '7' || b == '|' {
			hasUp = true
		}
	}

	if down, ok := map_.Down(startLoc); ok {
		b := map_.At(down)
		if b == 'L' || b == 'J' || b == '|' {
			hasDown = true
		}
	}

	if left, ok := map_.Left(startLoc); ok {
		b := map_.At(left)
		if b == 'L' || b == 'F' || b == '-' {
			hasLeft = true
		}
	}

	if right, ok := map_.Right(startLoc); ok {
		b := map_.At(right)
		if b == 'J' || b == '7' || b == '-' {
			hasRight = true
		}
	}

	var newByte byte
	switch {
	case hasUp && hasDown:
		newByte = '|'
	case hasLeft && hasRight:
		newByte = '-'
	case hasUp && hasLeft:
		newByte = 'J'
	case hasDown && hasRight:
		newByte = 'F'
	case hasDown && hasLeft:
		newByte = '7'
	case hasUp && hasRight:
		newByte = 'L'
	default:
		panic("unexpected starting location adjacent pipes")
	}

	// Set the byte back so our dfs algorithm can do it's thing.
	map_.bytes[startLoc] = newByte

	return MapResult{Map: &map_, Start: startLoc}
}

func (s Solution) Part1(input io.Reader) int {
	data := CreateMap(input)

	depths := maxDistanceBfs(data.Start, data.Map)

	SaveImageP1(data.Map, data.Start, depths)

	return maxDepth(depths)
}

type NodeDepth struct {
	Id    NodeId
	Depth int
}

type NodeDepthQueue struct {
	ids []NodeDepth
}

func (q *NodeDepthQueue) Push(n NodeDepth) {
	q.ids = append([]NodeDepth{n}, q.ids...)
}

func (q *NodeDepthQueue) Pop() (NodeDepth, bool) {
	if len(q.ids) == 0 {
		return NodeDepth{}, false
	}

	n := q.ids[len(q.ids)-1]
	q.ids = q.ids[0 : len(q.ids)-1]

	return n, true
}

func maxDistanceBfs(start NodeId, m *Map) map[NodeId]int {
	nodeDistance := make(map[NodeId]int)

	var queue NodeDepthQueue
	queue.Push(NodeDepth{Id: start, Depth: 0})

	for {
		this, ok := queue.Pop()
		if !ok {
			break
		}
		nodeDistance[this.Id] = this.Depth

		for _, conn := range m.ConnsAt(this.Id) {
			if _, ok := nodeDistance[conn]; ok {
				// we have already visited this node
			} else {
				queue.Push(NodeDepth{Id: conn, Depth: this.Depth + 1})
			}
		}
	}

	return nodeDistance
}

func maxDepth(in map[NodeId]int) int {
	m := 0
	for _, v := range in {
		if v > m {
			m = v
		}
	}
	return m
}

// Convert the Map from a map of pipes to a map that contains only our main loop, used for the algorithm in
// part 2.
func CreatePart2Map(in *Map, depths map[NodeId]int) *Map {
	var newBytes bytes.Buffer

	for idx := range in.bytes {
		b := in.bytes[idx]
		if b == '.' || b == '\n' {
			newBytes.WriteByte(b)
		} else if _, ok := depths[NodeId(idx)]; ok {
			newBytes.WriteByte(b)
		} else {
			newBytes.WriteByte('.')
		}
	}

	return &Map{
		bytes: newBytes.Bytes(),
		width: in.width,
	}
}

// Scan each line of the part 2 map and find areas inside the polygon using
// a similar approach to:
// https://www.tutorialspoint.com/computer_graphics/polygon_filling_algorithm.htm
func ScanPart2Map(m *Map) {
	idx := 0
	oddEven := 0

	// Are we in a vertical pipe segment? If so, defer odd/even calculation until
	// we see if it turns the same or opposite direction.
	inPipe := false
	var pipeStart byte

	for {
		b := m.bytes[idx]
		switch b {
		case '|':
			oddEven++
		case 'F', 'L':
			if inPipe {
				panic("unexpected vertical pipe start")
			}
			inPipe = true
			pipeStart = b

		case '7':
			if !inPipe {
				panic("expected vertical pipe state when turning out")
			} else if pipeStart == 'F' {
				oddEven += 2
			} else if pipeStart == 'L' {
				oddEven++
			} else {
				panic("unexpected byte " + string(b))
			}
			inPipe = false

		case 'J':
			if !inPipe {
				panic("expected vertical pipe state when turning out")
			} else if pipeStart == 'L' {
				oddEven += 2
			} else if pipeStart == 'F' {
				oddEven++
			} else {
				panic("unexpected byte " + string(b))
			}
			inPipe = false

		case '-', 'x':
			// nop

		case '.':
			if oddEven%2 == 1 {
				m.bytes[idx] = 'I'
			}

		case '\n':
			if oddEven%2 != 0 {
				panic("unexpected odd/even at end of line")
			}
			oddEven = 0
		}

		idx++
		if idx == len(m.bytes) {
			break
		}
	}
}

func (s Solution) Part2(input io.Reader) int {
	data := CreateMap(input)

	depths := maxDistanceBfs(data.Start, data.Map)

	newMap := CreatePart2Map(data.Map, depths)

	ScanPart2Map(newMap)

	SaveImageP2(newMap)

	return bytes.Count(newMap.bytes, []byte{'I'})
}
