package script

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
)

func TestN(t *testing.T) {
	n := NewN(mem.New())

	assert.Equal(t, 0, n.Run([]string{}))
	assert.Equal(t, 0, n.Run([]string{"2"}))
	assert.Equal(t, "\n\n\n", string(*n.Memory.Buffers[n.Memory.Destination]))
}
