package device

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/config"
	"github.com/thorntonrose/device/internal/mem"
)

func TestMain(m *testing.M) {
	defer config.InitLog()()
	m.Run()
}

//-----------------------------------------------------------------------------

func TestLoad(t *testing.T) {
	device := New()

	device.Load("003=FOO\n020$X")
	assert.Equal(t, []byte("FOO"), device.Memory.Slots[3])
	assert.Equal(t, []byte("X"), device.Memory.Slots[20])
}

func TestRun(t *testing.T) {
	device := New()
	device.Load("003=FOO\n020$X")

	device.Run(20)
	assert.Equal(t, []byte("FOO"), *device.Memory.Buffers[mem.Transmit])
}
