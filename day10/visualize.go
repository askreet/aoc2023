package day10

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

var (
	White = color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}
	Green = color.RGBA{0x00, 0xFF, 0x00, 0xFF}
	Red   = color.RGBA{0xFF, 0x00, 0x00, 0xFF}
	Black = color.Black
	Grey  = color.Gray{0x80}
)

func SaveImageP1(m *Map, startLoc NodeId, nodeDistance map[NodeId]int) {
	width := m.width - 1 // sans newline
	height := len(m.bytes) / m.width

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			point := image.Pt(x, y)

			if NodeAsXY(m, startLoc) == point {
				img.Set(x, y, White)
			} else if IsPartOfLoop(m, point, nodeDistance) {
				img.Set(x, y, Green)
			} else {
				img.Set(x, y, Black)
			}
		}
	}

	f, err := os.Create("day10_part1.png")
	if err != nil {
		panic(err)
	}
	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}

func SaveImageP2(m *Map) {
	width := m.width - 1 // sans newline
	height := len(m.bytes) / m.width

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			switch m.bytes[(y*m.width)+x] {
			case 'I':
				img.Set(x, y, Red)
			case 'x':
				img.Set(x, y, Grey)
			case '.':
				img.Set(x, y, Black)
			default:
				img.Set(x, y, Green)
			}

		}
	}

	f, err := os.Create("day10_part2.png")
	if err != nil {
		panic(err)
	}
	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}

func NodeAsXY(m *Map, n NodeId) image.Point {
	if n == 0 {
		return image.Pt(0, 0)
	}

	return image.Point{
		X: int(n) % m.width,
		Y: int(n) / m.width,
	}
}

func IsPartOfLoop(m *Map, p image.Point, loopNodes map[NodeId]int) bool {
	// This is expensive.
	for nodeId := range loopNodes {
		if NodeAsXY(m, nodeId) == p {
			return true
		}
	}

	return false
}
