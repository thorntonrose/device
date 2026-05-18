package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
)

func TestH(t *testing.T) {
	h := NewH(mem.New())
	*h.Memory.Buffers[h.Memory.Source] = []byte{'A', '\t', 'B'}

	AssertH(t, h, []string{}, 0, 0, 1)           // defaults
	AssertH(t, h, []string{"1"}, 0, 0, 1)        // '\t' found
	AssertH(t, h, []string{"1", "'A'"}, 0, 0, 0) // 'A' found
	AssertH(t, h, []string{"1", "'B'"}, 1, 0, 2) // 'B' found from position 1
	AssertH(t, h, []string{"1", "'C'"}, 1, 1, 1) // 'C' not found from position 1
}

func AssertH(t *testing.T, h H, parameters []string, ptr int, expectedSkip int, expectedPtr int) {
	h.Memory.Pointers[h.Memory.Source] = ptr

	skip := h.Run(parameters)
	assert.Equal(t, expectedSkip, skip)
	assert.Equal(t, expectedPtr, h.Memory.Pointers[h.Memory.Source])
}
