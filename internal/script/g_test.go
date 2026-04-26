package script

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
)

func TestG(t *testing.T) {
	memory := mem.New()
	*memory.Buffers[mem.Transmit] = []byte("WORLD")

	NewG(memory).Run([]string{})
	assert.Equal(t, []byte{}, *memory.Buffers[mem.Transmit])
}
