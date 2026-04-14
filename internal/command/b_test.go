package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
)

func TestB_NoChange(t *testing.T) {
	memory := mem.New()

	NewB(memory).Run([]string{})
	assert.Equal(t, mem.Receive, memory.Source)
	assert.Equal(t, mem.Transmit, memory.Destination)
}

func TestB_Set(t *testing.T) {
	memory := mem.New()

	NewB(memory).Run([]string{"1", "2"})
	assert.Equal(t, mem.Transmit, memory.Source)
	assert.Equal(t, mem.Receive, memory.Destination)
}

func TestB_Clear(t *testing.T) {
	memory := mem.New()
	memory.Slots[mem.Slot(mem.Receive)] = []byte("HELLO")
	memory.Get(mem.Pointers)[mem.Receive] = 5

	NewB(memory).Run([]string{"9"})
	assert.Equal(t, byte(0), memory.Get(mem.Pointers)[mem.Receive])
}
