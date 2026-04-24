package script

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
)

func TestX(t *testing.T) {
	memory := mem.New()
	memory.Set(mem.BufSlot(mem.Receive), []byte("HELLO"))

	NewX(memory).Run([]string{})
	assert.Equal(t, []byte("HELLO"), memory.Get(mem.BufSlot(mem.Transmit)))
}

// ???: Need more tests.
