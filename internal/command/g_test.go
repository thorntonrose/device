package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
)

func TestG(t *testing.T) {
	memory := mem.New()
	memory.Set(mem.Slot(mem.Transmit), []byte("HELLO"))

	NewG(memory).Run([]string{})
	assert.Equal(t, []byte{}, memory.Get(mem.Slot(mem.Transmit)))
}
