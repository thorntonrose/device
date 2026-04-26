package script

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
)

func TestV(t *testing.T) {
	v := NewV(mem.New())
	*v.Memory.Buffers[mem.Transmit] = []byte("FOO")
	*v.Memory.Buffers[mem.Receive] = []byte("BAR")

	assert.Equal(t, "FOO", v.Get([]string{}))
	assert.Equal(t, "BAR", v.Get([]string{"2"}))

	v.Run([]string{"2"})
}

func TestV_InvalidParameters(t *testing.T) {
	memory := mem.New()
	v := NewV(memory)

	assert.Panics(t, func() { v.Run([]string{"A"}) })
	assert.Panics(t, func() { v.Run([]string{"6"}) })
}
