package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
)

func TestAt(t *testing.T) {
	memory := mem.New()
	memory.Set(20, []byte("HELLO"))

	assert.Equal(t, 0, NewA(memory).Run([]string{"20", "0"}))
	assert.Equal(t, []byte("HELLO"), memory.Get(mem.Slot(mem.Transmit)))
}

func TestA_Skip(t *testing.T) {
	memory := mem.New()
	memory.Set(20, []byte{})

	assert.Equal(t, 1, NewA(memory).Run([]string{"20", "1"}))
}
