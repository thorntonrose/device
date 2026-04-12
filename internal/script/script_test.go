package script

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/command"
	"github.com/thorntonrose/device/internal/mem"
)

func TestRun(t *testing.T) {
	memory := mem.New()
	memory.Set(mem.Slot(mem.Receive), []byte("HELLO"))
	memory.Set(20, []byte("X"))

	New(&memory, map[string]command.Runner{"X": command.NewX(&memory)}).Run(20)
	assert.Equal(t, []byte("HELLO"), memory.Get(mem.Slot(mem.Transmit)))
}
