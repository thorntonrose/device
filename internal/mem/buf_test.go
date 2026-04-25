package mem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	memory := New()
	memory.Set(Receive+1, []byte{'A'})

	assert.Equal(t, byte('A'), memory.Read(Receive))
	assert.Equal(t, 1, memory.Pointers[Receive])
}

func TestReadAll(t *testing.T) {
	memory := New()
	memory.Set(Receive+1, []byte{'A', 'B', 'C'})

	assert.Equal(t, []byte{'A', 'B', 'C'}, memory.ReadAll(Receive, 0, 0))
}

func TestReadAll_Max(t *testing.T) {
	memory := New()
	memory.Set(Receive+1, []byte{'A', 'B', 'C'})

	assert.Equal(t, []byte{'A', 'B'}, memory.ReadAll(Receive, 2, 0))
}

func TestReadAll_Stop(t *testing.T) {
	memory := New()
	memory.Set(Receive+1, []byte{'A', 28, 'B'})

	assert.Equal(t, []byte{'A'}, memory.ReadAll(Receive, 0, 28))
}

func TestWriteAll(t *testing.T) {
	memory := New()

	memory.WriteAll(Transmit, []byte{'A'})
	assert.Equal(t, []byte{'A'}, *memory.Buffers[Transmit])

}

func TestClear(t *testing.T) {
	memory := New()
	memory.WriteAll(Transmit, []byte{'A'})

	memory.Clear(Transmit)
	assert.Equal(t, []byte{}, *memory.Buffers[Transmit])
}
