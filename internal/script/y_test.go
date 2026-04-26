package script

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/etc"
	"github.com/thorntonrose/device/internal/mem"
)

func TestY(t *testing.T) {
	c := NewY(mem.New())
	c.Memory.Slots[0] = []byte("FOO")
	c.Memory.Slots[1] = []byte("BAR")

	etc.Times(2, func(i int) {
		c.Memory.Clear(c.Memory.Destination)
		assert.Equal(t, 0, c.Run([]string{"1"}))
		assert.Equal(t, c.Format(i, c.Memory.Slots[i]), *c.Memory.Buffers[c.Memory.Destination])
	})

	c.SlotNum = len(c.Memory.Slots) - 1
	assert.Equal(t, 1, c.Run([]string{"1"}))
	assert.Equal(t, 0, c.SlotNum)
}
