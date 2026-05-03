package script

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thorntonrose/device/internal/mem"
)

func TestI(t *testing.T) {
	i := NewI(mem.New())
	*i.Memory.Buffers[mem.Receive] = []byte{28}

	assert.Equal(t, 0, i.Run([]string{}))               // nop
	assert.Equal(t, 1, i.Run([]string{"1"}))            // true
	assert.Equal(t, 1, i.Run([]string{"1", "1"}))       // equal
	assert.Equal(t, 1, i.Run([]string{"1", "2", "29"})) // not equal
	assert.Equal(t, 1, i.Run([]string{"1", "3", "29"})) // less than
	assert.Equal(t, 1, i.Run([]string{"1", "4", "27"})) // greater than

	assert.Panics(t, func() { i.Run([]string{"", "5"}) })
}

func TestI_String(t *testing.T) {
	i := NewI(mem.New())
	*i.Memory.Buffers[mem.Receive] = []byte("AB")

	assert.Equal(t, 1, i.Run([]string{"1", "1", "'AB'"}))
	assert.Equal(t, 0, i.Run([]string{"1", "1", "'ABC'"}))
}

func TestI_Empty(t *testing.T) {
	i := NewI(mem.New())
	assert.Equal(t, 0, i.Run([]string{"1", "1"}))
}
