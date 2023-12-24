package day22

import (
	"maps"
	"slices"
)

type Loc struct{ x, y, z int }
type Tower struct {
	shapes   []Shape
	shapeIdx map[Loc]int
}

// Settle will generate a new tower where all pieces have moved far down as possible without colliding with another piece.
func (t *Tower) Settle() *Tower {
	newTower := t.Clone()

	for newTower.SettleStep() > 0 {
	}

	return newTower
}

// Count the number of bricks that will settle, total.
func (t *Tower) SettleCount() int {
	for i := range t.shapes {
		t.shapes[i].bit = false
	}

	for {
		if t.SettleStep() == 0 {
			break
		}
	}

	sum := 0

	for i := range t.shapes {
		if t.shapes[i].bit {
			sum++
		}
	}

	return sum
}

// SettleStep executes a single settle action by mutating the current Tower. It returns the number of changed shapes.
func (t *Tower) SettleStep() int {
	numMoved := 0

	var lastSuccessfulLocation Shape
	var candidateLocation Shape
	for this := range t.shapes {
		if t.shapes[this].MinZ() > 0 {
			lastSuccessfulLocation = t.shapes[this]
			for {
				candidateLocation = lastSuccessfulLocation.Drop()
				collides := false
				// check the shape index
				candidateLocation.EachLoc(func(x, y, z int) bool {
					if idx, ok := t.shapeIdx[Loc{x, y, z}]; ok && idx != this {
						collides = true
						return false // do not keep searching
					}

					return true
				})

				//for other := range t.shapes {
				//	if this == other {
				//		continue
				//	}
				//
				//	if candidateLocation.Overlaps(t.shapes[other]) {
				//		collides = true
				//		break
				//	}
				//}
				if collides || lastSuccessfulLocation.MinZ() == 0 {
					t.RelocateIndex(this, lastSuccessfulLocation)
					break
				} else {
					lastSuccessfulLocation = candidateLocation
					numMoved++
				}
			}
		}
	}

	return numMoved
}

func (t *Tower) Clone() *Tower {
	return &Tower{
		shapes:   slices.Clone(t.shapes),
		shapeIdx: maps.Clone(t.shapeIdx),
	}
}

func (t *Tower) DeleteIndex(i int) {
	t.shapes = slices.Delete(t.shapes, i, i+1)

	for k := range t.shapeIdx {
		if v, ok := t.shapeIdx[k]; ok && v > i {
			t.shapeIdx[k] -= 1
		} else if v == i {
			delete(t.shapeIdx, k)
		}
	}
}

func (t *Tower) RelocateIndex(idx int, s Shape) {
	t.shapes[idx].EachLoc(func(x, y, z int) bool {
		delete(t.shapeIdx, Loc{x, y, z})
		return true
	})
	t.shapes[idx] = s
	s.EachLoc(func(x, y, z int) bool {
		t.shapeIdx[Loc{x, y, z}] = idx
		return true
	})

}

func (t *Tower) Add(shape Shape) {
	if t.shapeIdx == nil {
		t.shapeIdx = make(map[Loc]int)
	}

	t.shapes = append(t.shapes, shape)
	shape.EachLoc(func(x, y, z int) bool {
		t.shapeIdx[Loc{x, y, z}] = len(t.shapes)
		return true
	})
}

func (t *Tower) Len() int {
	return len(t.shapes)
}

func (t *Tower) Shapes() []Shape {
	return t.shapes
}
