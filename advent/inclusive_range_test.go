package advent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSparseRangeSet_Insert(t *testing.T) {
	var srs SparseRangeSet

	srs.Insert(IRange(0, 2))
	srs.Insert(IRange(4, 6))
	srs.Insert(IRange(2, 4))

	assert.Equal(t,
		[]InclusiveRange{IRange(0, 6)},
		srs.Ranges)
	assert.Equal(t, 7, srs.Total())
}
