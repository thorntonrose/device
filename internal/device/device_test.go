package device

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
)

func TestLoad(t *testing.T) {
	device := New()

	device.Load("003=HELLO\n020$X")
	assert.Equal(t, []byte("HELLO"), device.Memory.Slots[3])
	assert.Equal(t, []byte("X"), device.Memory.Slots[20])
}

func TestRun(t *testing.T) {
	device := New()
	device.Load("003=HELLO\n020$X")

	device.Run(20)
	assert.Equal(t, []byte("HELLO"), *device.Memory.Buffers[mem.Transmit])
}
