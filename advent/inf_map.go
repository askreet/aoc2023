package advent

import "fmt"

type infMapEntry struct {
	X int
	Y int
	B byte
}

type InfMap struct {
	coords []infMapEntry
}

func (i *InfMap) Set(x, y int, b byte) {
	i.coords = append(i.coords, infMapEntry{x, y, b})

}

// Create a bounded Map by determining the outer bounds of set pixels and filling the remainder in
// with b. The resulting map will use 0, 0 as an origin point.
func (i *InfMap) AsMap(b byte) *Map {
	top, left := 0, 0
	bottom, right := 0, 0
	for _, p := range i.coords {
		if p.X < left {
			left = p.X
		}
		if p.Y < top {
			top = p.Y
		}
		if p.X > right {
			right = p.X
		}
		if p.Y > bottom {
			bottom = p.Y
		}
	}

	fmt.Println(top, left, right, bottom)

	width := right - left + 1
	height := bottom - top + 1
	sz := (width + 1) * height
	fmt.Println("allocating", sz, "bytes of memory")
	var mapData []byte = make([]byte, sz)
	fmt.Println("done")

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			mapData = append(mapData, b)
		}
		mapData = append(mapData, '\n')
	}

	m := NewMap(mapData)
	for _, c := range i.coords {
		m.Set(c.X-left, c.Y-top, c.B)
	}

	return &m
}
