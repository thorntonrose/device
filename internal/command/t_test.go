package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
)

func TestT(t *testing.T) {
	tc := NewT(mem.New())

	assert.Equal(t, 0, tc.Run([]string{}))
	assert.Equal(t, []byte("1"), tc.Memory.Slots[0])

	assert.Panics(t, func() { tc.Run([]string{"", "5"}) })
}

func TestT_Subtract(t *testing.T) {
	tc := NewT(mem.New())
	tc.Memory.Slots[0] = []byte("1")

	assert.Equal(t, 0, tc.Run([]string{"0", "1"}))
	assert.Equal(t, []byte("0"), tc.Memory.Slots[0])
}

func TestT_Multiply(t *testing.T) {
	tc := NewT(mem.New())
	tc.Memory.Slots[0] = []byte("1")

	assert.Equal(t, 0, tc.Run([]string{"", "2", "2"}))
	assert.Equal(t, []byte("2"), tc.Memory.Slots[0])
}

func TestT_Divide(t *testing.T) {
	tc := NewT(mem.New())
	tc.Memory.Slots[0] = []byte("2")

	assert.Equal(t, 0, tc.Run([]string{"", "3", "2"}))
	assert.Equal(t, []byte("1"), tc.Memory.Slots[0])
}

func TestT_Modulo(t *testing.T) {
	tc := NewT(mem.New())
	tc.Memory.Slots[0] = []byte("5")

	assert.Equal(t, 0, tc.Run([]string{"", "4", "2"}))
	assert.Equal(t, []byte("1"), tc.Memory.Slots[0])
}

func TestT_Skip(t *testing.T) {
	tc := NewT(mem.New())
	tc.Memory.Slots[0] = []byte("1")

	assert.Equal(t, 1, tc.Run([]string{"", "1", "", "1"}))
}
