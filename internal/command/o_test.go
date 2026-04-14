package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
)

func TestO(t *testing.T) {
	memory := mem.New()
	memory.Set(mem.Slot(mem.Receive), []byte("HELLO"))

	o := NewO(memory)

	o.Run([]string{})
	assert.Equal(t, byte(1), memory.Get(mem.Pointers)[memory.Source])

	o.Run([]string{"1"})
	assert.Equal(t, byte(2), memory.Get(mem.Pointers)[memory.Source])
}
