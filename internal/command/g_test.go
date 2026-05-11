package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
	_ "github.com/thorntonrose/device/internal/testing"
)

func TestG(t *testing.T) {
	g := NewG(mem.New())
	*g.Memory.Buffers[mem.Transmit] = []byte("WORLD")

	g.Run([]string{})
	assert.Equal(t, []byte{}, *g.Memory.Buffers[mem.Transmit])
}
