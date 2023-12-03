package day2

type Bag struct {
	CubeCounts map[string]int
}

func NewBag() *Bag {
	return &Bag{
		CubeCounts: make(map[string]int),
	}
}

func (b *Bag) AddN(color string, n int) {
	b.CubeCounts[color] = n
}

func (b *Bag) Power() int {
	return b.CubeCounts["red"] * b.CubeCounts["green"] * b.CubeCounts["blue"]
}

func (b *Bag) CouldReveal(r *Reveal) bool {
	count, ok := b.CubeCounts[r.Color]
	if !ok {
		return false
	}

	return count >= r.Number
}

func MinBagFor(g *Game) *Bag {
	neededColors := make(map[string]int, 3)

	for _, set := range g.Sets {
		for _, reveal := range set {
			if v, ok := neededColors[reveal.Color]; ok {
				neededColors[reveal.Color] = max(reveal.Number, v)
			} else {
				neededColors[reveal.Color] = reveal.Number
			}
		}
	}

	return &Bag{
		CubeCounts: neededColors,
	}
}
