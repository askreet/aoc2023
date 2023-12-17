package advent

// Map is a 2D grid of bytes with newlines embedded within.
type Map struct {
	Bytes  []byte
	Width  int
	Height int
}

func NewMap(bs []byte) Map {
	var m Map

	var i = 0
	for {
		if bs[i] == '\n' {
			m.Bytes = bs
			m.Width = i
			m.Height = len(bs) / (m.Width + 1)
			return m
		}
		i++
	}
}

func (m *Map) At(x, y int) byte {
	idx := (y * (m.Width + 1)) + x
	return m.Bytes[idx]
}
