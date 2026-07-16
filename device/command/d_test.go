package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/device/mem"
	"github.com/thorntonrose/device/device/mem/buf"
)

func TestD(t *testing.T) {
	d := NewD(mem.New())
	buf.WriteAll(d.Memory, mem.Transmit, []byte{'A', 'B'})

	d.Run([]string{})
	assert.Equal(t, []byte{'A'}, *d.Memory.Buffers[mem.Transmit])

	assert.Panics(t, func() { d.Run([]string{"-1"}) })
}
