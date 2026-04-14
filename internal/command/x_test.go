package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
)

func TestX(t *testing.T) {
	memory := mem.New()
	memory.Set(mem.Slot(mem.Receive), []byte("HELLO"))

	NewX(memory).Run([]string{})
	assert.Equal(t, []byte("HELLO"), memory.Get(mem.Slot(mem.Transmit)))
}

// ???: Need more tests.
