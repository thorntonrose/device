package device

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/device/mem"
	_ "github.com/thorntonrose/device/device/testing"
)

func TestLoad(t *testing.T) {
	device := New()

	device.Load("003=FOO\n020$X")
	assert.Equal(t, []byte("FOO"), device.Memory.Slots[3])
	assert.Equal(t, []byte("X"), device.Memory.Slots[20])
}

func TestRun(t *testing.T) {
	device := New()
	device.Load("003=FOO\n020$X")

	device.Run()
	assert.Equal(t, []byte("FOO"), *device.Memory.Buffers[mem.Transmit])
}
