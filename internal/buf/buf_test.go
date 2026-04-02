package buf

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
)

func TestNewBufferSet(t *testing.T) {
	memory := mem.New()

	bufferSet := NewBufferSet(memory, 1)
	assert.Len(t, bufferSet.Buffers, MinBuffers+1)
	assert.Equal(t, &bufferSet.Buffers[Receive], bufferSet.Source)
	assert.Equal(t, &bufferSet.Buffers[Transmit], bufferSet.Destination)
}

func TestCopy(t *testing.T) {
	memory := mem.New()

	bufferSet := NewBufferSet(memory, 0)
	bufferSet.Source.Write('A')
	bufferSet.Source.Reset()

	bufferSet.Copy()
	assert.Equal(t, []byte{'A'}, bufferSet.Destination.Get())
	assert.Equal(t, []byte{'A'}, memory[mem.Transmit])
}

//-----------------------------------------------------------------------------

func TestNewBuffer(t *testing.T) {
	memory := mem.New()

	buffer := NewBuffer(memory, mem.Transmit)
	assert.NotNil(t, buffer.Memory)
	assert.Equal(t, mem.Transmit, buffer.Location)
}

func TestWrite(t *testing.T) {
	memory := mem.New()
	buffer := NewBuffer(memory, mem.Transmit)

	buffer.Write('A')
	assert.Equal(t, []byte{'A'}, buffer.Get())
	assert.Equal(t, 1, buffer.Pointer)
}

func TestRead(t *testing.T) {
	memory := mem.New()

	buffer := NewBuffer(memory, mem.Transmit)
	buffer.Write('A')
	buffer.Reset()

	data := buffer.Read()
	assert.Equal(t, byte('A'), data)
	assert.Equal(t, 1, buffer.Pointer)
}

func TestReset(t *testing.T) {
	buffer := NewBuffer(mem.New(), mem.Transmit)
	buffer.Pointer = 5

	buffer.Reset()
	assert.Equal(t, 0, buffer.Pointer)
}
