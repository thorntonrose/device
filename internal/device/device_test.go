package device

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
)

func TestLoad(t *testing.T) {
	device := New()

	device.Load("003=HELLO\n020$X")
	assert.Equal(t, []byte("HELLO"), device.Memory.Get(3))
	assert.Equal(t, []byte("X"), device.Memory.Get(20))
}

func TestRun(t *testing.T) {
	device := New()
	device.Load("003=HELLO\n020$X")

	device.Run(20)
	assert.Equal(t, []byte("HELLO"), device.Memory.Get(mem.BufSlot(mem.Transmit)))
}
