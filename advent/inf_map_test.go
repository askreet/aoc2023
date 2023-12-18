package advent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfMap_AsMap(t *testing.T) {
	var i InfMap

	i.Set(-1, -1, '#')
	i.Set(-0, -1, '#')
	i.Set(+1, -1, '#')

	i.Set(-1, +0, '#')

	i.Set(+1, +0, '#')

	i.Set(-1, +1, '#')
	i.Set(+0, +1, '#')
	i.Set(+1, +1, '#')

	m := i.AsMap('.')

	assert.Equal(t, 3, m.Width)
	assert.Equal(t, 3, m.Height)
	assert.Equal(t, "###\n#.#\n###\n", string(m.Bytes))
}
