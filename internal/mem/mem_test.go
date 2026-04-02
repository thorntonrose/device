package mem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	memory := New()
	assert.Len(t, memory, MaxLocations)
}

func TestSet(t *testing.T) {
	data := []byte("FOO")
	memory := New()

	memory.Set(20, data)
	assert.Equal(t, data, memory[20])
}

func TestAppend(t *testing.T) {
	memory := New()

	memory.Append(20, 'A')
	assert.Equal(t, []byte{'A'}, memory[20])
}
