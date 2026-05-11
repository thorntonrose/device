package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
	_ "github.com/thorntonrose/device/internal/testing"
)

func TestY(t *testing.T) {
	y := NewY(mem.New())
	y.Memory.Slots[0] = []byte("FOO")
	y.Memory.Slots[1] = []byte("BAR")
	y.Memory.Slots[3] = []byte("BAZ")

	for _, i := range []int{0, 1, 3} {
		y.Memory.Clear(y.Memory.Destination)
		assert.Equal(t, 0, y.Run([]string{"1"}))
		assert.Equal(t, y.Format(i, y.Memory.Slots[i]), *y.Memory.Buffers[y.Memory.Destination])
		assert.Equal(t, i+1, y.SlotNum)
	}

	assert.Equal(t, 1, y.Run([]string{"1"}))
	assert.Equal(t, 0, y.SlotNum)
}
