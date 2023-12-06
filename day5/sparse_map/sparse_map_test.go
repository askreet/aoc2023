package sparse_map

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	in := bytes.NewBufferString("49 53 8\n0 11 42\n42 0 7\n57 7 4\n")

	actual := Parse(in)

	expected := &Map{
		Entries: []Entry{
			{
				InStart:  49,
				OutStart: 53,
				Len:      8,
			},
			{
				InStart:  0,
				OutStart: 11,
				Len:      42,
			},
			{
				InStart:  42,
				OutStart: 0,
				Len:      7,
			},
			{
				InStart:  57,
				OutStart: 7,
				Len:      4,
			},
		},
	}

	assert.Equal(t, expected, actual)
}
