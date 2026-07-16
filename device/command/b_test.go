package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/device/mem"
	_ "github.com/thorntonrose/device/device/testing"
)

func TestB_NoChange(t *testing.T) {
	b := NewB(mem.New())
	b.Memory.Pointers[mem.Receive] = 1

	b.Run([]string{})
	assert.Equal(t, mem.Receive, b.Memory.Source)
	assert.Equal(t, mem.Transmit, b.Memory.Destination)
	assert.Equal(t, 1, b.Memory.Pointers[mem.Receive])

	assert.Panics(t, func() { b.Run([]string{"6"}) })
	assert.Panics(t, func() { b.Run([]string{"", "6"}) })
}

func TestB_Set(t *testing.T) {
	memory := mem.New()

	NewB(memory).Run([]string{"1", "2"})
	assert.Equal(t, 1, memory.Source)
	assert.Equal(t, 2, memory.Destination)
}

func TestB_Clear(t *testing.T) {
	memory := mem.New()
	memory.Slots[mem.Receive+1] = []byte("FOO")
	memory.Pointers[mem.Receive] = 5

	NewB(memory).Run([]string{"9"})
	assert.Equal(t, 0, memory.Pointers[mem.Receive])
}
