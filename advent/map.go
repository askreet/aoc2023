package advent

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

// Map is a 2D grid of bytes with newlines embedded within.
type Map struct {
	Bytes  []byte
	Width  int
	Height int
}

func NewMap(bs []byte) Map {
	var m Map

	if bs[len(bs)-1] != '\n' {
		bs = append(bs, '\n')
	}

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

func (m *Map) InBounds(x, y int) bool {
	return x >= 0 &&
		y >= 0 &&
		x < m.Width &&
		y < m.Height
}

func (m *Map) At(x, y int) byte {
	idx := (y * (m.Width + 1)) + x
	return m.Bytes[idx]
}

func (m *Map) Set(x int, y int, b byte) {
	idx := (y * (m.Width + 1)) + x
	m.Bytes[idx] = b
}

func (m *Map) Print(id int) {
	fmt.Printf("== Map %d ==\n", id)
	os.Stdout.Write(m.Bytes)
	//fmt.Print(m.Bytes)
	fmt.Println()
}

type VisitOrder struct {
	step  int
	xInit int
	yInit int
	xCond func(int) bool
	yCond func(int) bool
}

func (m *Map) OrderNormal() *VisitOrder {
	return &VisitOrder{
		step:  1,
		xInit: 0,
		yInit: 0,
		xCond: func(x int) bool { return x < m.Width },
		yCond: func(y int) bool { return y < m.Height },
	}
}

func (m *Map) OrderInverse() *VisitOrder {
	return &VisitOrder{
		step:  -1,
		xInit: m.Width - 1,
		yInit: m.Height - 1,
		xCond: func(x int) bool { return x >= 0 },
		yCond: func(y int) bool { return y >= 0 },
	}
}

func (m *Map) VisitAll(b byte, o *VisitOrder, f func(x int, y int)) {
	for y := o.yInit; o.yCond(y); y += o.step {
		for x := o.xInit; o.xCond(x); x += o.step {
			if m.At(x, y) == b {
				f(x, y)
			}
		}
	}
}

func (m *Map) Copy() *Map {
	var newMap Map
	newMap.Width = m.Width
	newMap.Height = m.Height
	newMap.Bytes = bytes.Clone(m.Bytes)
	return &newMap
}

func (m *Map) Count(b byte) int {
	return bytes.Count(m.Bytes, []byte{b})
}

func (m *Map) SavePNG(filename string, colormap map[byte]color.Color) {
	img := image.NewRGBA(image.Rect(0, 0, m.Width, m.Height))

	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			b := m.At(x, y)
			if c, ok := colormap[b]; ok {
				img.Set(x, y, c)
			} else {
				img.Set(x, y, color.Black)
			}
		}
	}

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}
