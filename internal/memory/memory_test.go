package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	data := []byte("FOO")
	memory := New()

	memory.Set(20, data)
	assert.Equal(t, data, memory.Locations[20])
}

func TestGet(t *testing.T) {
	data := []byte("FOO")
	memory := New()

	memory.Set(20, data)
	assert.Equal(t, data, memory.Get(20))
}

func TestDump(t *testing.T) {
	memory := New()

	for i := range memory.Locations {
		if !memory.IsReserved(i) {
			memory.Set(i, []byte("A"))
		}
	}

	memory.Dump()
}
