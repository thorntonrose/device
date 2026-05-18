package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
)

func TestD(t *testing.T) {
	d := NewD(mem.New())
	d.Memory.WriteAll(mem.Transmit, []byte{'A', 'B'})

	d.Run([]string{})
	assert.Equal(t, []byte{'A'}, *d.Memory.Buffers[mem.Transmit])

	assert.Panics(t, func() { d.Run([]string{"-1"}) })
}
