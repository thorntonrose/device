package script

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
)

func TestO(t *testing.T) {
	memory := mem.New()
	memory.Set(mem.Receive+1, []byte("FOO"))

	o := NewO(memory)

	o.Run([]string{})
	assert.Equal(t, 1, memory.Pointers[memory.Source])

	o.Run([]string{"1"})
	assert.Equal(t, 2, memory.Pointers[memory.Source])
}

func TestO_InvalidParameters(t *testing.T) {
	memory := mem.New()
	o := NewO(memory)

	assert.Panics(t, func() { o.Run([]string{"A"}) })
}
