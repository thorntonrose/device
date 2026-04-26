package script

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
)

func TestX(t *testing.T) {
	memory := mem.New()
	memory.Set(mem.Receive+1, []byte("FOO"))

	NewX(memory).Run([]string{})
	assert.Equal(t, []byte("FOO"), *memory.Buffers[mem.Transmit])
}

func TestX_InvalidParameters(t *testing.T) {
	memory := mem.New()
	memory.Set(mem.Receive+1, []byte("FOO"))
	x := NewX(memory)

	assert.Panics(t, func() { x.Run([]string{"A"}) })
	assert.Panics(t, func() { x.Run([]string{"0", "A"}) })
}
