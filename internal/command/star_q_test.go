package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
)

func TestStarQ(t *testing.T) {
	sq := NewStarQ(mem.New())
	AssertStarQ(t, sq, []string{}, 0, []byte("2748"), 2748)              // string (“2748” -> 2748)
	AssertStarQ(t, sq, []string{"#0", "1", "1"}, 0, []byte("V"), 86)     // ASCII (“V” -> 86)
	AssertStarQ(t, sq, []string{"#0", "1", "1"}, 0, []byte("XY"), 22617) // ASCII ("XY" -> 22617)

	assert.Panics(t, func() { sq.Run([]string{"#10"}) })
	assert.Panics(t, func() { sq.Run([]string{"", "2"}) })
}

func AssertStarQ(t *testing.T, sq StarQ, parameters []string, varNum int, data []byte, expected int) {
	*sq.Memory.Buffers[sq.Memory.Destination] = data
	sq.Memory.Variables[varNum] = 0

	sq.Run(parameters)
	assert.Equal(t, expected, sq.Memory.Variables[varNum])
}
