package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
)

func TestO(t *testing.T) {
	memory := mem.New()
	memory.Set(mem.BufSlot(mem.Receive), []byte("HELLO"))

	o := NewO(memory)

	o.Run([]string{})
	assert.Equal(t, 1, memory.Pointers[memory.Source])

	o.Run([]string{"1"})
	assert.Equal(t, 2, memory.Pointers[memory.Source])
}
