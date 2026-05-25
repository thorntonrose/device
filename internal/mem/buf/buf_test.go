package buf

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
)

func TestRead(t *testing.T) {
	memory := mem.New()
	memory.Set(mem.Receive+1, []byte{'A'})

	assert.Equal(t, byte('A'), Read(memory, mem.Receive))
	assert.Equal(t, 1, memory.Pointers[mem.Receive])
}

func TestReadAll(t *testing.T) {
	memory := mem.New()
	memory.Set(mem.Receive+1, []byte{'A', 'B', 'C'})

	assert.Equal(t, []byte{'A', 'B', 'C'}, ReadAll(memory, mem.Receive, 0, 0))
}

func TestReadAll_Max(t *testing.T) {
	memory := mem.New()
	memory.Set(mem.Receive+1, []byte{'A', 'B', 'C'})

	assert.Equal(t, []byte{'A', 'B'}, ReadAll(memory, mem.Receive, 2, 0))
}

func TestReadAll_Stop(t *testing.T) {
	memory := mem.New()
	memory.Set(mem.Receive+1, []byte{'A', '\n', 'B'})

	assert.Equal(t, []byte{'A'}, ReadAll(memory, mem.Receive, 0, '\n'))
}

func TestWriteAll(t *testing.T) {
	memory := mem.New()

	WriteAll(memory, mem.Transmit, []byte{'A'})
	assert.Equal(t, []byte{'A'}, *memory.Buffers[mem.Transmit])

}

func TestClear(t *testing.T) {
	memory := mem.New()
	WriteAll(memory, mem.Transmit, []byte{'A'})

	Clear(memory, mem.Transmit)
	assert.Equal(t, []byte{}, *memory.Buffers[mem.Transmit])
}
