package script

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
)

func TestX(t *testing.T) {
	x := NewX(mem.New())
	x.Memory.Set(mem.Receive+1, []byte("FOO"))

	x.Run([]string{})
	assert.Equal(t, []byte("FOO"), *x.Memory.Buffers[mem.Transmit])

	assert.Panics(t, func() { x.Run([]string{"1", "256"}) })
}
